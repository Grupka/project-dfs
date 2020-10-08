package naming_server

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"project-dfs/pb"
)

type NamingServerController struct {
	pb.UnimplementedNamingServer
	Server *NamingServer
}

//returns the pointer to the implementation
func NewNamingServiceController(server *NamingServer) *NamingServerController {
	return &NamingServerController{
		Server: server,
	}
}

func (ctlr *NamingServerController) Register(ctx context.Context, request *pb.RegRequest) (*pb.RegResponse, error) {
	// update address map on the NAMING Server

	otherPeer, ok := peer.FromContext(ctx)
	if !ok {
		println("other peer not found")
		return &pb.RegResponse{Status: pb.Status_DECLINE}, errors.New("other peer not found")
	}

	// add a new Server to the list of known Storage Servers
	peerAddress := otherPeer.Addr
	ctlr.Server.SetAddressMap(request.ServerAlias, peerAddress.String())

	// broadcast the address to all other Storage Servers
	for key, element := range ctlr.Server.StorageAddresses {
		if key == request.ServerAlias {
			continue
		}
		conn, err := grpc.Dial(element, grpc.WithInsecure())
		client := pb.NewStorageClient(conn)
		client.AddStorage(context.Background(),
			&pb.AddRequest{ServerAlias: request.ServerAlias, ServerAddress: peerAddress.String()})
		CheckError(err)
	}

	return &pb.RegResponse{Status: pb.Status_ACCEPT}, nil
}

func (ctlr *NamingServerController) Discover(ctx context.Context, request *pb.DiscoverRequest) (response *pb.DiscoverResponse, err error) {

	// key is the file's path
	// element is StorageInfo struct

	// if path == "" return ALL storage servers
	if request.Path == "" {
		storages := make([]*pb.DiscoveredStorage, 0)
		for alias, address := range ctlr.Server.StorageAddresses {
			storages = append(storages, &pb.DiscoveredStorage{
				Alias:   alias,
				Address: address,
			})
		}
		return &pb.DiscoverResponse{StorageInfo: storages}, nil
	}

	for key, element := range ctlr.Server.IndexMap {
		if key == request.Path {
			// if found file

			response = &pb.DiscoverResponse{}
			storages := make([]*pb.DiscoveredStorage, 0)
			for _, alias := range element.ServersList {
				storages = append(storages, &pb.DiscoveredStorage{
					Alias:   alias,
					Address: ctlr.Server.StorageAddresses[alias],
				})
			}

			response.StorageInfo = storages
			return response, nil
		}
	}
	return &pb.DiscoverResponse{StorageInfo: make([]*pb.DiscoveredStorage, 0)}, nil
}

// ---

func (ctlr *NamingServerController) CreateFile(ctx context.Context, request *pb.CreateFileRequest) (*pb.CreateFileResponse, error) {

}

func (ctlr *NamingServerController) Move(ctx context.Context, request *pb.MoveRequest) (*pb.MoveResponse, error) {

}

func (ctlr *NamingServerController) DeleteFile(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {

}

func (ctlr *NamingServerController) Copy(ctx context.Context, request *pb.CopyRequest) (*pb.CopyResponse, error) {
	panic("no copy operation")
}

func (ctlr *NamingServerController) DeleteDirectory(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {

}

func (ctlr *NamingServerController) MakeDirectory(ctx context.Context, request *pb.MakeDirectoryRequest) (*pb.MakeDirectoryResponse, error) {

}

func (ctlr *NamingServerController) ListDirectory(ctx context.Context, request *pb.ListDirectoryRequest) (*pb.ListDirectoryResponse, error) {

}
