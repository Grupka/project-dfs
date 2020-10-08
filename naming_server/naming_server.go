package naming_server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"project-dfs/pb"
	"sync"
)

type StorageInfo struct {
	ServersList []string // Server's aliases
}

type Node struct {
	Name     string // name of a file or directory
	Children []*Node
}

type NamingServer struct {
	storageAddressesMutex sync.Mutex
	StorageAddresses      map[string]string // key:value = serverAlias:serverAddress
	LocalAddress          string
	RootIndexNode         Node
}

func (server *NamingServer) SetAddressMap(newKey string, newValue string) {
	server.storageAddressesMutex.Lock()
	defer server.storageAddressesMutex.Unlock()
	server.StorageAddresses[newKey] = newValue
}

func initNamingServer() *NamingServer {
	// Obtain address from environment
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "0.0.0.0:5678"
		fmt.Println("ADDRESS variable not specified; falling back to", address)
	}

	rootNode := Node{
		Name:     "",
		Children: make([]*Node, 0),
	}

	return &NamingServer{
		storageAddressesMutex: sync.Mutex{},
		StorageAddresses:      make(map[string]string),
		LocalAddress:          address,
		RootIndexNode:         rootNode,
	}
}

func CheckError(err error) {
	if err != nil {
		println("Error serving gRPC naming server:", err.Error())
		os.Exit(1)
	}
}

func Run() {
	server := initNamingServer()

	println("Initialized metadata: ")
	fmt.Printf("%+v\n", server)

	listener, err := net.Listen("tcp", server.LocalAddress)
	CheckError(err)
	println("Listening on " + server.LocalAddress)

	namingController := NewNamingServiceController(server)
	grpcServer := grpc.NewServer()
	pb.RegisterNamingServer(grpcServer, namingController)
	err = grpcServer.Serve(listener)
	CheckError(err)
}
