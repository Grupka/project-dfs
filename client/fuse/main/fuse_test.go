package main

import (
	"fmt"
	"project-dfs/client"
	"project-dfs/client/fuse"
	"testing"
)

func TestMount(t *testing.T) {
	c := &client.Client{}

	root := fuse.NewDfsNode(c, "")
	server := fuse.Mount("mnt_point", root)

	fmt.Println("Waiting for unmount...")
	defer server.Unmount()
	server.Wait()
}
