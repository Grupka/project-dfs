package naming_server

import (
	"../pb"
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

//returns the pointer to the implementation
func NewRegistrationServiceController(metadataParam *NamingServer) *RegistrationServiceController {
	return &RegistrationServiceController{
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
	ctlr.metadata.SetAddressMap(request.ServerAlias, peerAddress.String())

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

type DiscoveryServiceController struct {
	pb.UnimplementedStorageDiscoveryServer
	metadata *NamingServer
}

func NewDiscoveryServiceController(metadataParam *NamingServer) *DiscoveryServiceController {
	return &DiscoveryServiceController{
		metadata: metadataParam,
	}
}

func (ctlr *DiscoveryServiceController) Discover(ctx context.Context, request *pb.DiscoverRequest) (response *pb.DiscoverResponse, err error) {

	for key, element := range ctlr.metadata.IndexMap {
		if key == request.Path {
			// if found file

			// element -> struct StorageInfo with map[alias]ip ServersList
			response = &pb.DiscoverResponse{}
			var resultMap []*pb.DiscoveredStorage
			for alias, ip := range element.ServersList {
				resultMap[alias] = ip
			}

			response.StorageInfo = resultMap
			return response, nil
		}
	}
	return &pb.DiscoverResponse{StorageInfo: nil}, nil
}
