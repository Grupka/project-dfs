package fuse

import (
	"context"
	"fmt"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"log"
	"syscall"
	"time"
)

var dirs = map[uint64]string{2: "dir1", 3: "dir2"}

type bytesFileHandle struct {
	content []byte
}

// bytesFileHandle allows reads
var _ = (fs.FileReader)((*bytesFileHandle)(nil))

//=== Files part ===//

func (*DfsNode) Open(ctx context.Context, flags uint32) (fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	fmt.Println("Open:", ctx, "; flags =", flags)

	now := time.Now().Format(time.StampNano) + "\n"
	fh = &bytesFileHandle{
		content: []byte(now),
	}

	// Return FOPEN_DIRECT_IO so content is not cached.
	return fh, fuse.FOPEN_DIRECT_IO, 0
}

func (fh *bytesFileHandle) Read(ctx context.Context, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	fmt.Println("Read:", fh, "; off =", off, "; ctx =", ctx)

	end := off + int64(len(dest))
	if end > int64(len(fh.content)) {
		end = int64(len(fh.content))
	}

	// We could copy to the `dest` buffer, but since we have a
	// []byte already, return that.
	return fuse.ReadResultData(fh.content[off:end]), 0
}

//=== Directories part ===//

func (node *DfsNode) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {
	fmt.Println("Readdir:", node.Inode.String())

	r := make([]fuse.DirEntry, 0, len(dirs))
	for ino, name := range dirs {
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
	fmt.Println("Lookup: name =", name, "; Parent node:", node, "; EntryOut:", out)

	// Check the list of files for the requested name
	ino := uint64(0)
	for _ino, _name := range dirs {
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
	operations := &DfsNode{}

	// The NewInode call wraps the `operations` object into an Inode.
	child := node.NewInode(ctx, operations, stable)

	// In case of concurrent lookup requests, it can happen that operations !=
	// child.Operations().
	return child, 0
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
