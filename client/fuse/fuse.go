package fuse

import (
	"fmt"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"log"
)

type DfsNode struct {
	fs.Inode
}

// Node types must be InodeEmbedders.
var _ = (fs.InodeEmbedder)((*DfsNode)(nil))

// Node types should implement some file system operations. Here we are testing it.
var _ = (fs.NodeLookuper)((*DfsNode)(nil))
var _ = (fs.NodeOpener)((*DfsNode)(nil))

func Mount(path string, node *DfsNode) *fuse.Server {
	fmt.Println("Mounting root to", path)

	server, err := fs.Mount(path, node, &fs.Options{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mounted successfully")
	return server
}
