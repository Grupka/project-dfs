package client

import (
	"../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

type Client struct {
	namingServerClient pb.NamingServerClient
}

func (client *Client) GetStorageServerForPath(path string) pb.FileOperationsManagerClient {
	return nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Errorf("error serving gRPC server %s", err)
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
