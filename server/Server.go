package main

import (
	"../pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Errorf("error serving gRPC server %s", err)
		os.Exit(1)
	}
}

func main() {
	port := ":5678"
	listener, err := net.Listen("tcp", port)
	CheckError(err)
	println("Listening on " + port)

	aliveServer := NewKeepAliveServer()
	grpcServer := grpc.NewServer()
	pb.RegisterKeepAliveServer(grpcServer, aliveServer)

	err = grpcServer.Serve(listener)
	CheckError(err)
}
