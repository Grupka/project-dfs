package main

import (
	"../pb"
	"context"
)

type RegistrationServiceController struct {
	//regService pb.UnimplementedRegistrationServer
	pb.UnimplementedRegistrationServer
}

//returns the pointer to the implementation
func NewRegistrationServiceController() *RegistrationServiceController {
	return &RegistrationServiceController{}
}

func (ctlr *RegistrationServiceController) Register(ctx context.Context, request *pb.RegRequest) (*pb.RegResponse, error) {
	// check whether storage server is in the same network
	/*if () {
		return &pb.RegResponse{Status: "ACCEPT"}, nil
	} else {
		return &pb.RegResponse{Status: "DECLINE"}, nil
	}*/

	// add address of a new server to the file with addresses
	//ctlr.regService.
	return &pb.RegResponse{Status: "ACCEPT"}, nil
}
