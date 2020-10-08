package client

import (
	"context"
	"google.golang.org/grpc"
	"math/rand"
	"project-dfs/pb"
)

type Client struct {
	NamingServerClient pb.NamingClient
	storageServers     map[string]pb.StorageClient
}

func (client *Client) GetStorageServerForPath(path string) pb.StorageClient {
	discoverResponse, err := client.NamingServerClient.Discover(context.Background(), &pb.DiscoverRequest{
		Path: path,
	})
	if err != nil {
		println("Error while getting storage server for path \""+path+"\":", err)
		return nil
	}
	if len(discoverResponse.StorageInfo) == 0 {
		println("No storage servers for path \"" + path + "\" found")
		return nil
	}
	randomIndex := rand.Intn(len(discoverResponse.StorageInfo))
	address := discoverResponse.StorageInfo[randomIndex].Address
	return client.GetStorageServerByAddress(address)
}

func (client *Client) GetStorageServerByAddress(address string) pb.StorageClient {
	server, ok := client.storageServers[address]
	if ok {
		// We are already connected to the server, just return the current connection
		return server
	}

	// Connect to the server and save it
	conn, err := grpc.Dial(address)
	if err != nil {
		println("Error occurred during connecting to storage server at \""+address+"\":", err)
		return nil
	}
	server = pb.NewStorageClient(conn)
	client.storageServers[address] = server

	return server
}

func (client *Client) GetRandomStorageServer() pb.StorageClient {
	return client.GetStorageServerByAddress("")
}

func Run() {
	namingServerAddress := "localhost:5678"

	_, err := grpc.Dial(namingServerAddress, grpc.WithInsecure())
	if err != nil {
		println("Error dialing naming server:", err.Error())
	}
}
