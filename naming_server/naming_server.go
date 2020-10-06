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
	}

	return &NamingServerMetadata{
		StorageAddresses: make(map[string]string),
		NetworkAddress:   "",
		Mask:             "",
		LocalAddress:     address,
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Errorf("error serving gRPC server %s", err)
		os.Exit(1)
	}
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
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
	println("successful serving")
	CheckError(err)

	println("Changed metadata: ")
	fmt.Printf("%+v\n", metadata)
}
