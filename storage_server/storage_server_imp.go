package storage_server

import (
	"context"
	"os"
	"project-dfs/pb"
	"syscall"
)

const (
	StoragePath = "storage"
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

type FileOperationServiceController struct {
	pb.UnimplementedFileOperationsManagerServer
	Server *StorageServer
}

func newFileOperationServiceController(server *StorageServer) *FileOperationServiceController {
	return &FileOperationServiceController{
		Server: server,
	}
}

func getFreeSpace() int64 {
	var stat syscall.Statfs_t
	wd, _ := os.Getwd()
	_ = syscall.Statfs(wd, &stat)
	return int64(stat.Bavail * uint64(stat.Bsize))
}

func (ctlr *FileOperationServiceController) Initialize(ctx context.Context, args *pb.InitializeArgs) (*pb.InitializeResult, error) {
	/* Initialize the client storage on a new system,
	remove any existing file in the dfs root directory and return available size.*/

	_ = os.RemoveAll(StoragePath)
	return &pb.InitializeResult{
		ErrorStatus: &pb.ErrorStatus{
			Code:        0,
			Description: "OK",
		},
		AvailableSize: getFreeSpace(),
	}, nil
}

func (ctlr *FileOperationServiceController) CreateFile(ctx context.Context, args *pb.CreateFileArgs) (*pb.CreateFileResult, error) {
	// create a new empty file

	path := StoragePath + args.Path
	_, err := os.Create(path)
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

func (ctlr *FileOperationServiceController) ReadFile(ctx context.Context, args *pb.ReadFileArgs) (response *pb.ReadFileResult, err error) {
	// download a file from the DFS to the Client side

	path := StoragePath + args.Path
	fd, err := os.Open(path)
	if err != nil {
		return &pb.ReadFileResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		},
			Buffer: make([]byte, 0),
			Count:  0}, nil
	}

	buf := make([]byte, args.Count)
	n, err := fd.ReadAt(buf, args.Offset)
	if err != nil {
		return &pb.ReadFileResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		},
			Buffer: make([]byte, 0),
			Count:  0}, nil
	}

	fd.Close()
	response = &pb.ReadFileResult{
		ErrorStatus: &pb.ErrorStatus{
			Code:        0,
			Description: "OK",
		},
		Buffer: buf[0:n],
		Count:  int32(n),
	}
	return response, nil
}

func (ctlr *FileOperationServiceController) WriteFile(ctx context.Context, args *pb.WriteFileArgs) (*pb.WriteFileResult, error) {

	path := StoragePath + args.Path
	fd, err := os.Open(path)
	if err != nil {
		return &pb.WriteFileResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		}}, nil
	}

	buf := args.Buffer
	_, err = fd.WriteAt(buf, args.Offset)
	if err != nil {
		return &pb.WriteFileResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		}}, nil
	}

	fd.Close()
	return &pb.WriteFileResult{ErrorStatus: &pb.ErrorStatus{
		Code:        0,
		Description: "OK",
	}}, nil
}

func (ctlr *FileOperationServiceController) Remove(ctx context.Context, args *pb.RemoveArgs) (*pb.RemoveResult, error) {
	// allow to delete any file from DFS
	// allow to delete directory.
	// If the directory contains files the system asks for confirmation

	path := StoragePath + args.Path
	err := os.RemoveAll(path)
	if err != nil {
		return &pb.RemoveResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		}}, nil
	}

	return &pb.RemoveResult{ErrorStatus: &pb.ErrorStatus{
		Code:        0,
		Description: "OK",
	}}, nil
}

func (ctlr *FileOperationServiceController) GetFileInfo(ctx context.Context, args *pb.GetFileInfoArgs) (*pb.GetFileInfoResult, error) {
	// provide information about the file (any useful information - size, node id, etc.)

	path := StoragePath + args.Path
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return &pb.GetFileInfoResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		},
			FileSize: 0}, nil
	}

	return &pb.GetFileInfoResult{ErrorStatus: &pb.ErrorStatus{
		Code:        0,
		Description: "OK",
	},
		FileSize: uint64(fileInfo.Size())}, nil
}

func (ctlr *FileOperationServiceController) Copy(ctx context.Context, args *pb.CopyArgs) (*pb.CopyResult, error) {

	path := StoragePath + args.Path

}

func (ctlr *FileOperationServiceController) Move(ctx context.Context, args *pb.MoveArgs) (*pb.MoveResult, error) {

	path := StoragePath + args.Path

}

func (ctlr *FileOperationServiceController) ReadDirectory(ctx context.Context, args *pb.ReadDirectoryArgs) (*pb.ReadDirectoryResult, error) {
	// return list of files, which are stored in the directory

	path := StoragePath + args.Path
	fd, err := os.Open(path)
	if err != nil {
		return &pb.ReadDirectoryResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		},
			Contents: make([]*pb.Node, 0)}, nil
	}

	fileInfo, err := fd.Readdir(0)
	if err != nil {
		return &pb.ReadDirectoryResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		},
			Contents: make([]*pb.Node, 0)}, nil
	}

	fileInfoEntries := make([]*pb.Node, len(fileInfo))
	var mode pb.NodeMode
	for _, entry := range fileInfo {
		if entry.IsDir() {
			mode = pb.NodeMode_DIRECTORY
		} else {
			mode = pb.NodeMode_REGULAR_FILE
		}
		fileInfoEntries = append(fileInfoEntries, &pb.Node{
			Mode: mode,
			Name: entry.Name(),
		})
	}

	return &pb.ReadDirectoryResult{ErrorStatus: &pb.ErrorStatus{
		Code:        0,
		Description: "OK",
	},
		Contents: fileInfoEntries}, nil
}

func (ctlr *FileOperationServiceController) MakeDirectory(ctx context.Context, args *pb.MakeDirectoryArgs) (*pb.MakeDirectoryResult, error) {

	path := StoragePath + args.Path
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return &pb.MakeDirectoryResult{ErrorStatus: &pb.ErrorStatus{
			Code:        1,
			Description: err.Error(),
		}}, nil
	}

	return &pb.MakeDirectoryResult{ErrorStatus: &pb.ErrorStatus{
		Code:        0,
		Description: "OK",
	}}, nil
}
