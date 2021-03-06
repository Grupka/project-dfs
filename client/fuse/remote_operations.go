package fuse

import (
	"context"
	"fmt"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"project-dfs/pb"
	"syscall"
	"unsafe"
)

type DfsHandle struct {
	Node *DfsNode
}

// Check that file handles implement read and write
// TODO: maybe remove this
var _ = (fs.FileReader)((*DfsHandle)(nil))
var _ = (fs.FileWriter)((*DfsHandle)(nil))

//=== General part ===//

// In our case, the most important things here are report file size and permissions (mode).
func (node *DfsNode) Getattr(ctx context.Context, f fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	path := node.Path()
	fmt.Println("Getattr: path:", path)

	out.Mode = 0777
	out.Size = 0

	// Find the appropriate storage server
	opClients := node.Client.GetStorageServersForPath(path)

	if len(opClients) == 0 {
		println("getattr: no storage servers found for", path)
		println("returning zero file size")

		return 0
		//return syscall.ENOENT
	}

	// Do the request to storage server
	info := pb.GetFileInfoArgs{
		Path: node.Path(),
	}

	for _, client := range opClients {
		result, err := client.GetFileInfo(ctx, &info)
		if err != nil {
			println("error occurred during getattr:", err.Error())
			continue
		}

		// Update the filesize
		out.Size = result.FileSize

		return 0
	}

	println("WARNING: completely failed getattr")
	return 1
}

func (node *DfsNode) Setattr(ctx context.Context, f fs.FileHandle, in *fuse.SetAttrIn, out *fuse.AttrOut) syscall.Errno {
	return 0
}

// Used to delete regular files.
func (node *DfsNode) Unlink(ctx context.Context, name string) syscall.Errno {
	path := node.PathForFile(name)
	fmt.Println("Unlink: path:", path)

	client := node.Client.NamingServerClient

	info := pb.DeleteRequest{
		Path: path,
	}

	result, err := client.DeleteFile(ctx, &info)
	if err != nil {
		println("error occurred during unlink:", err.Error())
		return syscall.EAGAIN
	}

	return syscall.Errno(result.ErrorStatus.Code)
}

// Used to delete directories.
func (node *DfsNode) Rmdir(ctx context.Context, name string) syscall.Errno {
	path := node.PathForFile(name)
	fmt.Println("Rmdir: path:", path)

	client := node.Client.NamingServerClient

	info := pb.DeleteRequest{
		Path: path,
	}

	result, err := client.DeleteDirectory(ctx, &info)
	if err != nil {
		println("error occurred during rmdir:", err.Error())
		return syscall.EAGAIN
	}

	return syscall.Errno(result.ErrorStatus.Code)
}

//=== Files part ===//

// Creates a handle for a file.
func (node *DfsNode) Open(ctx context.Context, flags uint32) (fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	fh = &DfsHandle{Node: node}

	// Return FOPEN_DIRECT_IO so content is not cached.
	return fh, fuse.FOPEN_DIRECT_IO, 0
}

// Reads from a file handle. In our case, simply forwards call to the node.
func (h *DfsHandle) Read(ctx context.Context, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	return h.Node.Read(ctx, h, dest, off)
}

// Reads from the node.
func (node *DfsNode) Read(ctx context.Context, f fs.FileHandle, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	path := node.Path()
	fmt.Println("Read: path:", path)

	// Find the appropriate storage server
	opClients := node.Client.GetStorageServersForPath(path)
	if len(opClients) == 0 {
		println("read: no storage servers found for", path)
		return nil, syscall.ENOENT
	}

	info := pb.ReadFileArgs{
		Path:   path,
		Offset: off,
		Count:  int64(len(dest)),
	}

	for _, client := range opClients {
		result, err := client.ReadFile(ctx, &info)
		if err != nil {
			println("error occurred during read:", err.Error())
			continue
		}

		println("Read bytes:", result.Buffer)
		println("Storage server response:", result.ErrorStatus.String())

		return fuse.ReadResultData(result.Buffer), syscall.Errno(result.ErrorStatus.Code)
	}

	println("WARNING: completely failed read")
	return nil, 1
}

// Writes to a handle. In our case, simply forwards call to the node.
func (h *DfsHandle) Write(ctx context.Context, data []byte, off int64) (written uint32, errno syscall.Errno) {
	return h.Node.Write(ctx, h, data, off)
}

// Writes to a node.
func (node *DfsNode) Write(ctx context.Context, f fs.FileHandle, data []byte, off int64) (written uint32, errno syscall.Errno) {
	path := node.Path()
	fmt.Println("Write: path:", path)

	// Find the appropriate storage server
	opClients := node.Client.GetStorageServersForPath(path)

	if len(opClients) == 0 {
		println("write: no storage server found for", path)
		return 0, syscall.ENOENT
	}

	info := pb.WriteFileArgs{
		Path:        path,
		Offset:      off,
		Buffer:      data,
		IsChainCall: false,
	}

	for _, client := range opClients {
		result, err := client.WriteFile(ctx, &info)
		if err != nil {
			println("error occurred during write:", err.Error())
			continue
		}

		println("Storage server response:", result.ErrorStatus.String())

		return uint32(len(data)), syscall.Errno(result.ErrorStatus.Code)
	}

	println("WARNING: completely failed write")
	return 0, 1
}

//=== Directories part ===//

// Lists all nodes in a directory.
func (node *DfsNode) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {
	path := node.Path()
	fmt.Println("Readdir: path:", path)

	// Find the appropriate storage server
	opClient := node.Client.NamingServerClient

	info := pb.ListDirectoryRequest{
		Path: path,
	}

	result, err := opClient.ListDirectory(ctx, &info)
	if err != nil {
		println("error occurred during readdir:", err.Error())
		return nil, syscall.EAGAIN
	}

	r := make([]fuse.DirEntry, 0, len(result.Contents))
	for _, n := range result.Contents {
		mode := fuse.S_IFREG
		if n.Mode == pb.NodeMode_DIRECTORY {
			mode = fuse.S_IFDIR
		}

		r = append(r, fuse.DirEntry{
			Name: n.Name,
			Mode: uint32(mode),
		})
	}

	return fs.NewListDirStream(r), syscall.Errno(result.ErrorStatus.Code)
}

// Checks if asked file is located in the asked node.
func (node *DfsNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	fmt.Println("lookup:", node, ";", name)

	path := node.Path()

	out.Mode = 0777

	// Find the appropriate storage server
	opClient := node.Client.NamingServerClient

	info := pb.ListDirectoryRequest{
		Path: path,
	}

	result, err := opClient.ListDirectory(ctx, &info)
	if err != nil {
		println("error occurred during lookup:", err.Error())
		return nil, syscall.EAGAIN
	}

	for _, n := range result.Contents {
		if n.Name != name {
			continue
		}

		mode := fuse.S_IFREG
		if n.Mode == pb.NodeMode_DIRECTORY {
			mode = fuse.S_IFDIR
		}

		operations := NewDfsNode(node.Client, name)
		stable := fs.StableAttr{Mode: uint32(mode)}
		child := node.NewInode(ctx, operations, stable)

		return child.EmbeddedInode(), 0
	}

	return nil, syscall.ENOENT
}

// Creates a file.
func (node *DfsNode) Create(ctx context.Context, name string, flags uint32, mode uint32, out *fuse.EntryOut) (n *fs.Inode, fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	path := node.PathForFile(name)
	fmt.Println("Create: path:", path, "; name:", name, "; nodepath:", node.Path())

	opClient := node.Client.NamingServerClient

	info := pb.CreateFileRequest{
		Path: path,
	}

	result, err := opClient.CreateFile(ctx, &info)
	if err != nil {
		println("error occurred during create:", err.Error())
		return nil, nil, 0, syscall.EAGAIN
	}

	operations := NewDfsNode(node.Client, name)
	stable := fs.StableAttr{Mode: fuse.S_IFREG}
	child := node.NewInode(ctx, operations, stable)

	return child, fh, 0, syscall.Errno(result.ErrorStatus.Code)
}

// Renames a node (both files and directories).
func (node *DfsNode) Rename(ctx context.Context, name string, newParent fs.InodeEmbedder, newName string, flags uint32) syscall.Errno {
	path := node.PathForFile(name)
	_newParent := (*DfsNode)(unsafe.Pointer(newParent.EmbeddedInode()))
	newPath := _newParent.PathForFile(newName)
	fmt.Println("Rename: oldPath:", path, "; newPath:", newPath)

	// Find the appropriate storage server
	opClient := node.Client.NamingServerClient

	info := pb.MoveRequest{
		Path:    path,
		NewPath: newPath,
	}

	result, err := opClient.Move(ctx, &info)
	if err != nil {
		println("error occurred during rename:", err.Error())
		return syscall.EAGAIN
	}

	return syscall.Errno(result.ErrorStatus.Code)
}

// Creates a directory.
func (node *DfsNode) Mkdir(ctx context.Context, name string, mode uint32, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	path := node.PathForFile(name)
	fmt.Println("Mkdir: path:", path)

	opClient := node.Client.NamingServerClient

	info := pb.MakeDirectoryRequest{
		Path: path,
	}

	result, err := opClient.MakeDirectory(ctx, &info)
	if err != nil {
		println("error occurred during mkdir:", err.Error())
		return nil, syscall.EAGAIN
	}

	operations := NewDfsNode(node.Client, name)
	stable := fs.StableAttr{Mode: fuse.S_IFDIR}
	child := node.NewInode(ctx, operations, stable)

	return child, syscall.Errno(result.ErrorStatus.Code)
}

//=== Locks part ===//

//func (*DfsNode) Getlk(ctx context.Context, owner uint64, lk *fuse.FileLock, flags uint32, out *fuse.FileLock) syscall.Errno {
//	log.Fatal("Getlk is not implemented")
//	return 0
//}
//
//func (*DfsNode) Setlk(ctx context.Context, f fs.FileHandle, owner uint64, lk *fuse.FileLock, flags uint32) syscall.Errno {
//	log.Fatal("Setlk is not implemented")
//	return 0
//}
//
//func (*DfsNode) Setlkw(ctx context.Context, owner uint64, lk *fuse.FileLock, flags uint32) syscall.Errno {
//	log.Fatal("Setlkw is not implemented")
//	return 0
//}
