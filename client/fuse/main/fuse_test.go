package main

import (
	"fmt"
	"project-dfs/client/fuse"
	"testing"
)

func TestMount(t *testing.T) {
	root := fuse.NewDfsNode("", []byte("hello"), map[uint64]string{2: "dir1", 3: "dir2"})
	server := fuse.Mount("mnt_point", &root)

	fmt.Println("Waiting for unmount...")
	defer server.Unmount()
	server.Wait()
}
