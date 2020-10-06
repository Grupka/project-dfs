package main

import (
	"../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

type StorageServerMetadata struct {
	LocalAddress        string
	Alias               string
	NamingServerAddress string
	StorageAddresses    map[string]string // key:value = serverAlias:serverAddress
}

func initMetadata() *StorageServerMetadata {

	// Obtain local address from environment
	localAddress := os.Getenv("ADDRESS")
	if localAddress == "" {
		localAddress = "0.0.0.0:5678"
		fmt.Println("ADDRESS variable not specified; falling back to", localAddress)
	}

	// Obtain naming address from environment
	namingServerAddress := os.Getenv("NAMING_SERVER_ADDRESS")
	if namingServerAddress == "" {
		namingServerAddress = "localhost:5678"
		fmt.Println("NAMING_SERVER_ADDRESS variable not specified; falling back to", namingServerAddress)
	}

	// Obtain alias from environment
	alias := os.Getenv("ALIAS")
	if alias == "" {
		alias = "storage"
		fmt.Println("ALIAS variable not specified; falling back to", alias)
	}

	return &StorageServerMetadata{
		LocalAddress:        localAddress,
		Alias:               alias,
		NamingServerAddress: namingServerAddress,
		StorageAddresses:    make(map[string]string),
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Errorf("error serving gRPC server %s", err)
		os.Exit(1)
	}
}

func main() {

	metadata := initMetadata()

	fmt.Printf("Initialized storage metadata: %+v\n", metadata)

	fmt.Println("Connecting to naming server at", metadata.NamingServerAddress)
	conn, err := grpc.Dial(metadata.NamingServerAddress, grpc.WithInsecure())
	CheckError(err)

	newServer := pb.NewRegistrationClient(conn)
	response, err := newServer.Register(context.Background(),
		&pb.RegRequest{ServerAlias: metadata.Alias})
	CheckError(err)
	log.Printf("Response from naming server: %s", response.GetStatus())
}
