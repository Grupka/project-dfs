package storage_server

import (
	"../pb"
	"context"
)

type AdditionServiceController struct {
	pb.UnimplementedStorageAdditionServer
	metadata *StorageServer
}

func NewAdditionServiceController(metadataParam *StorageServer) *AdditionServiceController {
	return &AdditionServiceController{
		metadata: metadataParam,
	}
}

func (ctlr *AdditionServiceController) AddStorage(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	// update address map on the STORAGE server
	ctlr.metadata.SetMap(request.GetServerAlias(), request.GetServerAddress())

	return &pb.AddResponse{Status: pb.Status_ACCEPT}, nil
}
