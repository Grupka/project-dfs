package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"os"
	"project-dfs/pb"
)

type Client struct {
	NamingServerClient pb.StorageDiscoveryClient
	storageServers     map[string]pb.FileOperationsManagerClient
}

func (client *Client) GetStorageServerForPath(path string) pb.FileOperationsManagerClient {
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

func (client *Client) GetStorageServerByAddress(address string) pb.FileOperationsManagerClient {
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
	server = pb.NewFileOperationsManagerClient(conn)
	client.storageServers[address] = server

	return server
}

func (client *Client) GetRandomStorageServer() pb.FileOperationsManagerClient {
	return client.GetStorageServerByAddress("")
}

func CheckError(err error) {
	if err != nil {
		println("Error serving gRPC server:", err)
		os.Exit(1)
	}
}

func Run() {
	namingServerAddress := "localhost:5678"

	conn, err := grpc.Dial(namingServerAddress, grpc.WithInsecure())
	CheckError(err)

	aliveClient := pb.NewKeepAliveClient(conn)

	response, err := aliveClient.Check(context.Background(), &pb.KeepAliveRequest{Message: "Hello From Client!"})
	CheckError(err)
	log.Printf("Response from server: %s", response.GetMessage())
}
