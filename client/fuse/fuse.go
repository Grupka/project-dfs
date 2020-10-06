package fuse

import (
	"fmt"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"log"
)

type DfsNode struct {
	fs.Inode
	Path     string
	Content  []byte
	Children map[string]*DfsNode
}

func NewDfsNode(path string, content []byte, children map[string]*DfsNode) *DfsNode {
	return &DfsNode{
		Path:     path,
		Content:  content,
		Children: children,
	}
}

func (node *DfsNode) String() string {
	return "DfsNode{" + node.Path + "}"
}

// Node types must be InodeEmbedders.
var _ = (fs.InodeEmbedder)((*DfsNode)(nil))

// Node types should implement some file system operations. Here we are testing it.
var _ = (fs.NodeLookuper)((*DfsNode)(nil))
var _ = (fs.NodeOpener)((*DfsNode)(nil))
var _ = (fs.NodeGetattrer)((*DfsNode)(nil))
var _ = (fs.NodeCreater)((*DfsNode)(nil))
var _ = (fs.NodeReader)((*DfsNode)(nil))
var _ = (fs.NodeWriter)((*DfsNode)(nil))
var _ = (fs.NodeUnlinker)((*DfsNode)(nil))
var _ = (fs.NodeRmdirer)((*DfsNode)(nil))

func Mount(path string, node *DfsNode) *fuse.Server {
	fmt.Println("Mounting root to", path)

	server, err := fs.Mount(path, node, &fs.Options{
		MountOptions: fuse.MountOptions{Debug: true},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mounted successfully")
	return server
}
