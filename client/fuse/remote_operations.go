package fuse

import (
	"context"
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
	out.Mode = 0777

	// Find the appropriate storage server
	opClient := node.Client.GetStorageServerForPath(node.Path())

	// Do the request to storage server
	info := pb.GetFileInfoArgs{
		Path: node.Path(),
	}
	result, err := opClient.GetFileInfo(ctx, &info)
	if err != nil {
		println("error occurred during getattr:", err)
		return syscall.EAGAIN
	}

	// Update the filesize
	out.Size = result.FileSize

	return 0
}

// Used to delete regular files.
func (node *DfsNode) Unlink(ctx context.Context, name string) syscall.Errno {
	path := node.Path() + "/" + name

	// Find the appropriate storage server
	opClient := node.Client.GetStorageServerForPath(path)

	info := pb.RemoveArgs{
		Path: path,
	}

	result, err := opClient.Remove(ctx, &info)
	if err != nil {
		println("error occurred during unlink:", err)
		return syscall.EAGAIN
	}

	return syscall.Errno(result.ErrorStatus.Code)
}

// Used to delete directories.
func (node *DfsNode) Rmdir(ctx context.Context, name string) syscall.Errno {
	path := node.Path() + "/" + name

	// Find the appropriate storage server
	opClient := node.Client.GetStorageServerForPath(path)

	info := pb.RemoveArgs{
		Path: path,
	}

	result, err := opClient.Remove(ctx, &info)
	if err != nil {
		println("error occurred during rmdir:", err)
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

	// Find the appropriate storage server
	opClient := node.Client.GetStorageServerForPath(path)

	info := pb.ReadFileArgs{
		Path:   path,
		Offset: off,
		Count:  int64(len(dest)),
	}

	result, err := opClient.ReadFile(ctx, &info)
	if err != nil {
		println("error occurred during read:", err)
		return nil, syscall.EAGAIN
	}

	return fuse.ReadResultData(result.Buffer), syscall.Errno(result.ErrorStatus.Code)
}

// Writes to a handle. In our case, simply forwards call to the node.
func (h *DfsHandle) Write(ctx context.Context, data []byte, off int64) (written uint32, errno syscall.Errno) {
	return h.Node.Write(ctx, h, data, off)
}

// Writes to a node.
func (node *DfsNode) Write(ctx context.Context, f fs.FileHandle, data []byte, off int64) (written uint32, errno syscall.Errno) {
	path := node.Path()

	// Find the appropriate storage server
	opClient := node.Client.GetStorageServerForPath(path)

	info := pb.WriteFileArgs{
		Path:   path,
		Offset: off,
		// TODO: remove count, as it is derivable from the buffer
		Count:  int64(len(data)),
		Buffer: data,
	}

	result, err := opClient.WriteFile(ctx, &info)
	if err != nil {
		println("error occurred during write:", err)
		return 0, syscall.EAGAIN
	}

	return uint32(len(data)), syscall.Errno(result.ErrorStatus.Code)
}

//=== Directories part ===//

// Lists all nodes in a directory.
func (node *DfsNode) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {
	path := node.Path()

	// Find the appropriate storage server
	opClient := node.Client.GetStorageServerForPath(path)

	info := pb.ReadDirectoryArgs{
		Path: path,
	}

	result, err := opClient.ReadDirectory(ctx, &info)
	if err != nil {
		println("error occurred during rmdir:", err)
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
	path := node.Path()

	// Find the appropriate storage server
	opClient := node.Client.GetStorageServerForPath(path)

	info := pb.ReadDirectoryArgs{
		Path: path,
	}

	result, err := opClient.ReadDirectory(ctx, &info)
	if err != nil {
		println("error occurred during lookup:", err)
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

	return nil, syscall.Errno(result.ErrorStatus.Code)
}

// Creates a file.
func (node *DfsNode) Create(ctx context.Context, name string, flags uint32, mode uint32, out *fuse.EntryOut) (n *fs.Inode, fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	path := node.Path() + "/" + name

	opClient := node.Client.GetRandomStorageServer()

	info := pb.CreateFileArgs{
		Path: path,
	}

	result, err := opClient.CreateFile(ctx, &info)
	if err != nil {
		println("error occurred during create:", err)
		return nil, nil, 0, syscall.EAGAIN
	}

	operations := NewDfsNode(node.Client, name)
	stable := fs.StableAttr{Mode: fuse.S_IFREG}
	child := node.NewInode(ctx, operations, stable)

	return child, fh, 0, syscall.Errno(result.ErrorStatus.Code)
}

// Renames a node (both files and directories).
func (node *DfsNode) Rename(ctx context.Context, name string, newParent fs.InodeEmbedder, newName string, flags uint32) syscall.Errno {
	path := node.Path() + "/" + name
	_newParent := (*DfsNode)(unsafe.Pointer(newParent.EmbeddedInode()))

	// Find the appropriate storage server
	opClient := node.Client.GetStorageServerForPath(path)

	info := pb.MoveArgs{
		Path:    path,
		NewPath: _newParent.Path() + "/" + newName,
	}

	result, err := opClient.Move(ctx, &info)
	if err != nil {
		println("error occurred during rename:", err)
		return syscall.EAGAIN
	}

	return syscall.Errno(result.ErrorStatus.Code)
}

// Creates a directory.
func (node *DfsNode) Mkdir(ctx context.Context, name string, mode uint32, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	path := node.Path()

	opClient := node.Client.GetRandomStorageServer()

	info := pb.MakeDirectoryArgs{
		Path: path,
	}

	result, err := opClient.MakeDirectory(ctx, &info)
	if err != nil {
		println("error occurred during create:", err)
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
