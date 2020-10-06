package naming_server

import (
	"../pb"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/peer"
)

type RegistrationServiceController struct {
	pb.UnimplementedRegistrationServer
	metadata *NamingServerMetadata
}

//returns the pointer to the implementation
func NewRegistrationServiceController(metadataParam *NamingServerMetadata) *RegistrationServiceController {
	return &RegistrationServiceController{
		metadata: metadataParam,
	}
}

func (ctlr *RegistrationServiceController) Register(ctx context.Context, request *pb.RegRequest) (*pb.RegResponse, error) {
	// add address of a new server to address map from metadata on the naming server

	otherPeer, ok := peer.FromContext(ctx)
	if !ok {
		fmt.Errorf("other peer not found")
		return &pb.RegResponse{Status: pb.Status_DECLINE}, errors.New("other peer not found")
	}

	peerAddress := otherPeer.Addr
	ctlr.metadata.StorageAddresses[request.ServerAlias] = peerAddress.String()
	return &pb.RegResponse{Status: pb.Status_ACCEPT}, nil
}
