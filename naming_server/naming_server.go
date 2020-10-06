package naming_server

import (
	"../pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

type NamingServerMetadata struct {
	StorageAddresses map[string]string // key:value = serverAlias:serverAddress
	NetworkAddress   string            // to store network ip + mask
	Mask             string
	LocalAddress     string
}

func initMetadata() *NamingServerMetadata {
	// Obtain address from environment
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "0.0.0.0:5678"
		fmt.Println("ADDRESS variable not specified; falling back to", address)
	}

	networkAddress := ""
	mask := ""

	return &NamingServerMetadata{
		StorageAddresses: make(map[string]string),
		NetworkAddress:   networkAddress,
		Mask:             mask,
		LocalAddress:     address,
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Errorf("error serving gRPC server %s", err)
		os.Exit(1)
	}
}

func Run() {
	metadata := initMetadata()

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
