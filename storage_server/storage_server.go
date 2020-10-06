package main

import (
	"../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Errorf("error serving gRPC server %s", err)
		os.Exit(1)
	}
}

func main() {
	// Obtain address from environment
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "0.0.0.0:5678"
		fmt.Println("ADDRESS variable not specified; falling back to", address)
	}

	// Obtain naming address from environment
	namingServerAddress := os.Getenv("NAMING_SERVER_ADDRESS")
	if namingServerAddress == "" {
		namingServerAddress = "localhost:5678"
		fmt.Println("NAMING_SERVER_ADDRESS variable not specified; falling back to", namingServerAddress)
	}

	fmt.Println("Connecting to naming server at", namingServerAddress)
	conn, err := grpc.Dial(namingServerAddress, grpc.WithInsecure())
	CheckError(err)

	newServer := pb.NewRegistrationClient(conn)
	response, err := newServer.Register(context.Background(), &pb.RegRequest{ServerAddress: address, ServerAlias: "storage01"})
	CheckError(err)
	log.Printf("Response from naming server: %s", response.GetStatus())
}
