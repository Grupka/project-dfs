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
	serverAddr := "localhost:5678"

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	CheckError(err)

	aliveClient := pb.NewKeepAliveClient(conn)

	response, err := aliveClient.Check(context.Background(), &pb.KeepAliveRequest{Message: "Hello From Client!"})
	CheckError(err)
	log.Printf("Response from server: %s", response.GetMessage())
}
