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
	ServersList []string // server's aliases
}

type NamingServer struct {
	storageAddressesMutex sync.Mutex
	StorageAddresses      map[string]string // key:value = serverAlias:serverAddress
	LocalAddress          string
	IndexMap              map[string]StorageInfo //  key:value = path:storage info
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

	return &NamingServer{
		storageAddressesMutex: sync.Mutex{},
		StorageAddresses:      make(map[string]string),
		LocalAddress:          address,
		IndexMap:              make(map[string]StorageInfo),
	}
}

func CheckError(err error) {
	if err != nil {
		println("Error serving gRPC server:", err)
		os.Exit(1)
	}
}

func Run() {
	metadata := initNamingServer()

	println("Initialized metadata: ")
	fmt.Printf("%+v\n", metadata)

	listener, err := net.Listen("tcp", metadata.LocalAddress)
	CheckError(err)
	println("Listening on " + metadata.LocalAddress)

	regController := NewRegistrationServiceController(metadata)
	grpcServer := grpc.NewServer()
	pb.RegisterRegistrationServer(grpcServer, regController)
	err = grpcServer.Serve(listener)
	CheckError(err)
}
