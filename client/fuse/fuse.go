package fuse

import (
	"fmt"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"log"
	"project-dfs/client"
	"strings"
	"unsafe"
)

type DfsNode struct {
	fs.Inode
	Client *client.Client
	Name   string
}

func NewDfsNode(client *client.Client, name string) *DfsNode {
	return &DfsNode{
		Client: client,
		Name:   name,
	}
}

func (node *DfsNode) PathForFile(name string) string {
	var nodes []string
	n := node

	for n != nil {
		nodes = append([]string{n.Name}, nodes...)
		_, n_ := n.Parent()
		n = (*DfsNode)(unsafe.Pointer(n_))
	}

	nodes = append(nodes, name)

	if len(nodes) == 0 {
		return "/"
	} else {
		return strings.Join(nodes, "/")
	}
}

func (node *DfsNode) Path() string {
	var nodes []string
	n := node

	for n != nil {
		nodes = append([]string{n.Name}, nodes...)
		_, n_ := n.Parent()
		n = (*DfsNode)(unsafe.Pointer(n_))
	}

	if len(nodes) == 0 {
		return "/"
	} else {
		return strings.Join(nodes, "/")
	}

	//n := node
	//path := ""
	//
	//// Handle root node
	//if n.Name == "" {
	//	return ""
	//}
	//
	//for n != nil {
	//	path = "/" + n.Name + path
	//	_, n_ := n.Parent()
	//	n = (*DfsNode)(unsafe.Pointer(n_))
	//}
	//
	//return path
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
		log.Fatal(err.Error())
	}

	fmt.Println("Mounted successfully")
	return server
}
