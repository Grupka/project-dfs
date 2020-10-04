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
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	//arguments := os.Args
	//serverAddr := arguments[1] // "host:port" as a string
	serverAddr := "localhost:5678"

	//var conn *grpc.ClientConn
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	CheckError(err)

	aliveClient := pb.NewKeepAliveClient(conn)

	response, err := aliveClient.Check(context.Background(), &pb.KeepAliveRequest{Message: "Hello From Client!"})
	CheckError(err)
	log.Printf("Response from server: %s", response.GetMessage())
}
