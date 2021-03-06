package main

import (
	"fmt"
	"google.golang.org/grpc"
	"os"
	"project-dfs/client"
	"project-dfs/client/fuse"
)

func main() {
	namingServerAddress, ok := os.LookupEnv("NAMING_SERVER_ADDRESS")
	if !ok {
		namingServerAddress = "localhost:5678"
		fmt.Println("NAMING_SERVER_ADDRESS not specified, falling back to", namingServerAddress)
	}

	conn, err := grpc.Dial(namingServerAddress, grpc.WithInsecure())
	if err != nil {
		println("Error occurred while connecting to naming server:", err.Error())
		return
	}

	c := client.NewClient(conn)

	root := fuse.NewDfsNode(c, "")
	server := fuse.Mount("mnt_point", root)

	fmt.Println("Waiting for unmount...")
	defer server.Unmount()
	server.Wait()
}
