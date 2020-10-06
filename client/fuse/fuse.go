package fuse

import (
	"fmt"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"log"
	"project-dfs/client"
	"unsafe"
)

type DfsNode struct {
	fs.Inode
	Client   *client.Client
	Name     string
	Content  []byte
	Children map[string]*DfsNode
}

func NewDfsNode(name string, content []byte, children map[string]*DfsNode) *DfsNode {
	return &DfsNode{
		Name:     name,
		Content:  content,
		Children: children,
	}
}

func (node *DfsNode) Path() string {
	n := node
	path := ""

	for n != nil {
		path = "/" + n.Name + path
		_, n_ := n.Parent()
		n = (*DfsNode)(unsafe.Pointer(n_))
	}

	return path
}

func (node *DfsNode) String() string {
	return "DfsNode{" + node.Path() + "}"
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
