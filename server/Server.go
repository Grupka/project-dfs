package main

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"project-dfs/pb"
)

func CheckError(err error) {
	if err != nil {
		println("Error serving gRPC server", err)
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
