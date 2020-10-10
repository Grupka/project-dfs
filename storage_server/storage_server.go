package storage_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"project-dfs/pb"
	"strconv"
	"strings"
	"sync"
)

type StorageServer struct {
	LocalAddress          string
	Alias                 string
	NamingServerAddress   string
	storageAddressesMutex sync.Mutex
	storageAddresses      map[string]string // key:value = serverAlias:serverAddress
	namingClient          pb.NamingClient
	storageClients        map[string]pb.StorageClient
}

func (server *StorageServer) SetMap(newKey string, newValue string) {
	server.storageAddressesMutex.Lock()
	defer server.storageAddressesMutex.Unlock()
	server.storageAddresses[newKey] = newValue
}

func (server *StorageServer) GetNamingClient() pb.NamingClient {
	if server.namingClient == nil {
		conn, err := grpc.Dial(server.NamingServerAddress, grpc.WithInsecure())
		if err != nil {
			println("Error while getting naming client:", err)
			return nil
		}
		server.namingClient = pb.NewNamingClient(conn)
	}
	return server.namingClient
}

func (server *StorageServer) GetStorageClient(address string) pb.StorageClient {
	client, ok := server.storageClients[address]
	if !ok {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			println("Error while getting storage client:", err)
			return nil
		}
		client = pb.NewStorageClient(conn)
		server.storageClients[address] = client
	}
	return client
}

func initStorageServer() *StorageServer {
	// Obtain local address from environment
	localAddress := os.Getenv("ADDRESS")
	if localAddress == "" {
		localAddress = "0.0.0.0:5678"
		fmt.Println("ADDRESS variable not specified; falling back to", localAddress)
	}

	// Obtain naming address from environment
	namingServerAddress := os.Getenv("NAMING_SERVER_ADDRESS")
	if namingServerAddress == "" {
		namingServerAddress = "localhost:5678"
		fmt.Println("NAMING_SERVER_ADDRESS variable not specified; falling back to", namingServerAddress)
	}

	// Obtain alias from environment
	alias := os.Getenv("ALIAS")
	if alias == "" {
		alias = "storage"
		fmt.Println("ALIAS variable not specified; falling back to", alias)
	}

	return &StorageServer{
		LocalAddress:          localAddress,
		Alias:                 alias,
		NamingServerAddress:   namingServerAddress,
		storageAddressesMutex: sync.Mutex{},
		storageAddresses:      make(map[string]string),
		storageClients:        map[string]pb.StorageClient{},
	}
}

func CheckError(err error) {
	if err != nil {
		println("error serving gRPC storage server:", err.Error())
		os.Exit(1)
	}
}

func Run() {
	metadata := initStorageServer()

	fmt.Printf("Initialized storage metadata: %+v\n", metadata)

	fmt.Println("Connecting to naming server at", metadata.NamingServerAddress)
	conn, err := grpc.Dial(metadata.NamingServerAddress, grpc.WithInsecure())
	CheckError(err)

	port, _ := strconv.Atoi(metadata.LocalAddress[strings.LastIndex(metadata.LocalAddress, ":")+1:])

	newServer := pb.NewNamingClient(conn)
	response, err := newServer.Register(context.Background(),
		&pb.RegRequest{ServerAlias: metadata.Alias, Port: uint32(port)})
	CheckError(err)
	log.Printf("Response from naming server: %s", response.GetStatus())

	if response.GetStatus().String() == "ACCEPT" {
		// listen to connections
		listener, err := net.Listen("tcp", metadata.LocalAddress)
		CheckError(err)
		println("Listening on " + metadata.LocalAddress)

		storageController := NewStorageServiceController(metadata)
		grpcServer := grpc.NewServer()
		pb.RegisterStorageServer(grpcServer, storageController)
		err = grpcServer.Serve(listener)
		CheckError(err)
	}

}
