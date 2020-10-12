package client

import (
	"context"
	"google.golang.org/grpc"
	"project-dfs/pb"
)

type Client struct {
	NamingServerClient pb.NamingClient
	storageServers     map[string]pb.StorageClient
}

func NewClient(conn grpc.ClientConnInterface) *Client {
	return &Client{
		NamingServerClient: pb.NewNamingClient(conn),
		storageServers:     map[string]pb.StorageClient{},
	}
}

func (client *Client) GetStorageServersForPath(path string) []pb.StorageClient {
	discoverResponse, err := client.NamingServerClient.Discover(context.Background(), &pb.DiscoverRequest{
		Path: path,
	})
	if err != nil {
		println("Error while getting storage server for path \""+path+"\":", err.Error())
		return nil
	}
	if len(discoverResponse.StorageInfo) == 0 {
		println("No storage servers for path \"" + path + "\" found")
		return nil
	}
	//randomIndex := rand.Intn(len(discoverResponse.StorageInfo))
	//address := discoverResponse.StorageInfo[randomIndex].PublicAddress
	var servers []pb.StorageClient
	for _, storage := range discoverResponse.StorageInfo {
		servers = append(servers, client.GetStorageServerByAddress(storage.PublicAddress))
	}
	return servers
}

func (client *Client) GetStorageServerByAddress(address string) pb.StorageClient {
	server, ok := client.storageServers[address]
	if ok {
		// We are already connected to the server, just return the current connection
		return server
	}

	// Connect to the server and save it
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		println("Error occurred during connecting to storage server at \""+address+"\":", err.Error())
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
