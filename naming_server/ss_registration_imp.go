package main

import (
	"../pb"
	"context"
)

type RegistrationServiceController struct {
	pb.UnimplementedRegistrationServer
	metadata *NamingServerMetadata // pointer, right?
}

//returns the pointer to the implementation
func NewRegistrationServiceController(metadataParam *NamingServerMetadata) *RegistrationServiceController {
	return &RegistrationServiceController{
		metadata: metadataParam,
	}
}

func (ctlr *RegistrationServiceController) Register(ctx context.Context, request *pb.RegRequest) (*pb.RegResponse, error) {
	// add address of a new server to address map from metadata on the naming server
	ctlr.metadata.StorageAddresses["server"] = request.ServerAddress
	return &pb.RegResponse{Status: pb.Status_ACCEPT}, nil
}
