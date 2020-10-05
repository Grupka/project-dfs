package main

import (
	"../pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

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

func main() {

	ip := GetLocalIP()
	port := "5678"
	localAddress := ip + ":" + port
	listener, err := net.Listen("tcp", ":"+port)
	CheckError(err)
	println("Listening on " + port)
	println("Local address " + localAddress)

	regController := NewRegistrationServiceController()
	grpcServer := grpc.NewServer()
	pb.RegisterRegistrationServer(grpcServer, regController)
	err = grpcServer.Serve(listener)
	CheckError(err)
}
