package fuse

import (
	"context"
	"fmt"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"log"
	"syscall"
	"time"
	"unsafe"
)

type bytesFileHandle struct {
	Node *DfsNode
}

//=== General part ===//

func (node *DfsNode) Getattr(ctx context.Context, f fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	out.Size = 9999
	out.Mode = 0777
	return 0
}

//=== Files part ===//

func (node *DfsNode) Open(ctx context.Context, flags uint32) (fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	//fmt.Println("Open:", node, "; flags =", flags)

	fh = &bytesFileHandle{Node: node}

	// Return FOPEN_DIRECT_IO so content is not cached.
	return fh, fuse.FOPEN_DIRECT_IO, 0
}

func (h *bytesFileHandle) Read(ctx context.Context, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	return h.Node.Read(ctx, dest, off)
}

func (node *DfsNode) Read(ctx context.Context, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	//fmt.Println("Read:", fh, "; off =", off, "; ctx =", ctx)

	end := off + int64(len(dest))
	if end > int64(len(node.Content)) {
		end = int64(len(node.Content))
	}

	// We could copy to the `dest` buffer, but since we have a
	// []byte already, return that.
	return fuse.ReadResultData(node.Content[off:end]), 0
}

func (h *bytesFileHandle) Write(ctx context.Context, data []byte, off int64) (written uint32, errno syscall.Errno) {
	return h.Node.Write(ctx, data, off)
}

func (node *DfsNode) Write(ctx context.Context, data []byte, off int64) (written uint32, errno syscall.Errno) {
	//fmt.Println("Write:", fh, "; off =", off, "; ctx =", ctx, "data =", data)

	// Extend the content size if needed
	end := off + int64(len(data))
	if end > int64(len(node.Content)) {
		additionalArray := make([]byte, end-int64(len(node.Content)))
		node.Content = append(node.Content, additionalArray...)
	}

	copy(node.Content[off:end], data)
	return uint32(len(data)), 0
}

//=== Directories part ===//

func (node *DfsNode) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {
	//fmt.Println("Readdir:", node.Inode.String())

	r := make([]fuse.DirEntry, 0, len(node.Children))
	for ino, name := range node.Children {
		d := fuse.DirEntry{
			Name: name,
			Ino:  ino,
			// In our FS, mode is either DIRectory (fuse.S_IFDIR) or REGular file (fuse.S_IFREG).
			// We do not support symlinks or other stuff.
			Mode: fuse.S_IFREG,
		}
		r = append(r, d)
	}

	fmt.Println("Readdir: returned", len(r), "entries")
	return fs.NewListDirStream(r), 0
}

func (node *DfsNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	// Check the list of files for the requested name
	ino := uint64(0)
	for _ino, _name := range node.Children {
		if _name == name {
			ino = _ino
			break
		}
	}

	// If no such entry is found in the directory, abort
	if ino == 0 {
		return nil, syscall.ENOENT
	}

	stable := fs.StableAttr{
		Mode: fuse.S_IFREG,
		// The child inode is identified by its Inode number.
		// If multiple concurrent lookups try to find the same
		// inode, they are deduplicated on this key.
		Ino: ino,
	}
	operations := &DfsNode{Path: node.Path + "/" + name}

	// The NewInode call wraps the `operations` object into an Inode.
	child := node.NewInode(ctx, operations, stable)

	// In case of concurrent lookup requests, it can happen that operations !=
	// child.Operations().
	return child, 0
}

func (node *DfsNode) Create(ctx context.Context, name string, flags uint32, mode uint32, out *fuse.EntryOut) (n *fs.Inode, fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	path := node.Path + "/" + name
	ino := FilepathHash(path)

	stable := fs.StableAttr{
		Mode: fuse.S_IFREG,
		// The child inode is identified by its Inode number.
		// If multiple concurrent lookups try to find the same
		// inode, they are deduplicated on this key.
		Ino: ino,
	}

	now := time.Now().Format(time.StampNano) + "\n"

	operations := NewDfsNode(path, []byte(now), map[uint64]string{})
	child := node.NewInode(ctx, &operations, stable)

	node.Children[ino] = name

	return child, fh, 0, 0
}

func (node *DfsNode) Rename(ctx context.Context, name string, newParent fs.InodeEmbedder, newName string, flags uint32) syscall.Errno {
	newNodeNode := newParent.EmbeddedInode()
	// TODO: replace with something more elegant?
	newNode := *(*DfsNode)(unsafe.Pointer(newNodeNode))

	newHash := FilepathHash(newName)
	hash := FilepathHash(name)

	//newNode.AddChild(newName, node.GetChild(name), true)
	node.MvChild(name, newParent.EmbeddedInode(), newName, true)

	newNode.Children[newHash] = node.Children[hash]
	delete(node.Children, hash)

	return 0
}

//=== Locks part ===//

func (*DfsNode) Getlk(ctx context.Context, owner uint64, lk *fuse.FileLock, flags uint32, out *fuse.FileLock) syscall.Errno {
	log.Fatal("Getlk is not implemented")
	return 0
}

func (*DfsNode) Setlk(ctx context.Context, f fs.FileHandle, owner uint64, lk *fuse.FileLock, flags uint32) syscall.Errno {
	log.Fatal("Setlk is not implemented")
	return 0
}

func (*DfsNode) Setlkw(ctx context.Context, owner uint64, lk *fuse.FileLock, flags uint32) syscall.Errno {
	log.Fatal("Setlkw is not implemented")
	return 0
}
