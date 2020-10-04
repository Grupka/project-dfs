package main

import (
	"fmt"
	"project-dfs/client/fuse"
	"testing"
)

func TestMount(t *testing.T) {
	root := &fuse.DfsNode{}
	server := fuse.Mount("mnt_point", root)

	fmt.Println("Waiting for unmount...")
	defer server.Unmount()
	server.Wait()
}
