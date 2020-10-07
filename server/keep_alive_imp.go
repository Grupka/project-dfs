package main

import (
	"context"
	"log"
	"project-dfs/pb"
)

//implementation of KeepAliveService
type KeepAliveServerImpl struct {
	pb.UnimplementedKeepAliveServer
}

//returns the pointer to the implementation
func NewKeepAliveServer() *KeepAliveServerImpl {
	return &KeepAliveServerImpl{}
}

//function for implementation of gRPC Service
func (s *KeepAliveServerImpl) Check(ctx context.Context, in *pb.KeepAliveRequest) (*pb.KeepAliveResponse, error) {
	log.Printf("Receive message body from client: %s", in.GetMessage())
	return &pb.KeepAliveResponse{Message: "i am alive!"}, nil
}
