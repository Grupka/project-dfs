package storage_server

import (
	"../pb"
	"context"
	"os"
	"syscall"
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
	/* Initialize the client storage on a new system,
	remove any existing file in the dfs root directory and return available size.*/
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

func CreateFile(ctx context.Context, args *pb.CreateFileArgs) (*pb.CreateFileResult, error) {

	//create a new empty file
	_, err := os.Create(args.Path)
	if err != nil {
		return &pb.CreateFileResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		}}, nil
	}

	return &pb.CreateFileResult{ErrorStatus: &pb.ErrorStatus{
		Code:        0,
		Description: "OK",
	}}, nil
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

	// download a file from the DFS to the Client side

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

	// upload a file from the Client side to the DFS
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

	// allow to delete any file from DFS
	// allow to delete directory.
	//If the directory contains files the system asks for confirmation

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

	// provide information about the file (any useful information - size, node id, etc.)
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

	// allow to create a copy of file

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

	// allow to move a file to the specified path
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

	// return list of files, which are stored in the directory
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

	// allow to create a new directory
}
