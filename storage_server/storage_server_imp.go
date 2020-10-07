package storage_server

import (
	"../pb"
	"context"
)

type AdditionServiceController struct {
	pb.UnimplementedStorageAdditionServer
	Server *StorageServer
}

func NewAdditionServiceController(server *StorageServer) *AdditionServiceController {
	return &AdditionServiceController{
		Server: server,
	}
}

func (ctlr *AdditionServiceController) AddStorage(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	// update address map on the STORAGE server
	ctlr.Server.SetMap(request.GetServerAlias(), request.GetServerAddress())

	return &pb.AddResponse{}, nil
}

// ---

type InitServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewInitServiceController(server *StorageServer) *InitServiceController {
	return &InitServiceController{
		Server: server,
	}
}

func Initialize(context.Context, *pb.InitializeArgs) (*pb.InitializeResult, error) {

}

// ---

type CreateFileServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewCreateFileServiceController(server *StorageServer) *CreateFileServiceController {
	return &CreateFileServiceController{
		Server: server,
	}
}

func CreateFile(context.Context, *pb.CreateFileArgs) (*pb.CreateFileResult, error) {

}

// ---

type ReadFileServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewReadFileServiceController(server *StorageServer) *ReadFileServiceController {
	return &ReadFileServiceController{
		Server: server,
	}
}

func ReadFile(context.Context, *pb.ReadFileArgs) (*pb.ReadFileResult, error) {

}

// ---

type WriteFileServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewWriteFileServiceController(server *StorageServer) *WriteFileServiceController {
	return &WriteFileServiceController{
		Server: server,
	}
}

func WriteFile(context.Context, *pb.WriteFileArgs) (*pb.WriteFileResult, error) {

}

// ---

type RemoveServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewRemoveServiceController(server *StorageServer) *RemoveServiceController {
	return &RemoveServiceController{
		Server: server,
	}
}

func Remove(context.Context, *pb.RemoveArgs) (*pb.RemoveResult, error) {

}

// ---

type GetFileInfoServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewGetFileInfoServiceController(server *StorageServer) *GetFileInfoServiceController {
	return &GetFileInfoServiceController{
		Server: server,
	}
}

func GetFileInfo(context.Context, *pb.GetFileInfoArgs) (*pb.GetFileInfoResult, error) {

}

// ---

type CopyServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewCopyServiceController(server *StorageServer) *CopyServiceController {
	return &CopyServiceController{
		Server: server,
	}
}

func Copy(context.Context, *pb.CopyArgs) (*pb.CopyResult, error) {

}

// ---

type MoveServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewMoveServiceController(server *StorageServer) *MoveServiceController {
	return &MoveServiceController{
		Server: server,
	}
}

func Move(context.Context, *pb.MoveArgs) (*pb.MoveResult, error) {

}

// ---

type ReadDirectoryServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewReadDirectoryServiceController(server *StorageServer) *ReadDirectoryServiceController {
	return &ReadDirectoryServiceController{
		Server: server,
	}
}

func ReadDirectory(context.Context, *pb.ReadDirectoryArgs) (*pb.ReadDirectoryResult, error) {

}

// ---

type MakeDirectoryServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func NewMakeDirectoryServiceController(server *StorageServer) *MakeDirectoryServiceController {
	return &MakeDirectoryServiceController{
		Server: server,
	}
}

func MakeDirectory(context.Context, *pb.MakeDirectoryArgs) (*pb.MakeDirectoryResult, error) {

}
