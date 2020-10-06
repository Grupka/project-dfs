package naming_server

import (
	"../pb"
	"../storage_server"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type RegistrationServiceController struct {
	pb.UnimplementedRegistrationServer
	metadata *NamingServer
}

type AdditionServiceController struct {
	pb.UnimplementedStorageAdditionServer
	metadata *storage_server.StorageServer
}

//returns the pointer to the implementation
func NewRegistrationServiceController(metadataParam *NamingServer) *RegistrationServiceController {
	return &RegistrationServiceController{
		metadata: metadataParam,
	}
}

func NewAdditionServiceController(metadataParam *storage_server.StorageServer) *AdditionServiceController {
	return &AdditionServiceController{
		metadata: metadataParam,
	}
}

func (ctlr *RegistrationServiceController) Register(ctx context.Context, request *pb.RegRequest) (*pb.RegResponse, error) {
	// update address map on the NAMING server

	otherPeer, ok := peer.FromContext(ctx)
	if !ok {
		fmt.Errorf("other peer not found")
		return &pb.RegResponse{Status: pb.Status_DECLINE}, errors.New("other peer not found")
	}

	// add a new server to the list of known Storage Servers
	peerAddress := otherPeer.Addr
	ctlr.metadata.SetMap(request.ServerAlias, peerAddress.String())

	// broadcast the address to all other Storage Servers
	for key, element := range ctlr.metadata.StorageAddresses {
		if key == request.ServerAlias {
			continue
		}
		conn, err := grpc.Dial(element, grpc.WithInsecure())
		client := pb.NewStorageAdditionClient(conn)
		client.AddStorage(context.Background(),
			&pb.AddRequest{ServerAlias: request.ServerAlias, ServerAddress: peerAddress.String()})
		CheckError(err)
	}

	return &pb.RegResponse{Status: pb.Status_ACCEPT}, nil
}

func (ctlr *AdditionServiceController) AddStorage(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	// update address map on the STORAGE server
	ctlr.metadata.SetMap(request.GetServerAlias(), request.GetServerAddress())

	return &pb.AddResponse{Status: pb.Status_ACCEPT}, nil
}
