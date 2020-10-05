package main

import (
	"../pb"
	"context"
)

type RegistrationServiceImpl struct {
	pb.UnimplementedRegistrationServer
}

//returns the pointer to the implementation
func NewRegistrationService() *RegistrationServiceImpl {
	return &RegistrationServiceImpl{}
}

func (s *RegistrationServiceImpl) Register(context.Context, *pb.RegRequest) (*pb.RegResponse, error) {
	// check whether storage server is in the same network
	return &pb.RegResponse{Status: "ACCEPT"}, nil
	/*if () {
		return &pb.RegResponse{Status: "ACCEPT"}, nil
	} else {
		return &pb.RegResponse{Status: "DECLINE"}, nil
	}*/
}
