package main

import (
	"../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
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

	//namingServerAddr := "10.5.0.254:5678" // CHANGE TO SOME ADDRESS
	namingServerAddr := "localhost:5678"

	conn, err := grpc.Dial(namingServerAddr, grpc.WithInsecure())
	CheckError(err)

	newServer := pb.NewRegistrationClient(conn)
	response, err := newServer.Register(context.Background(), &pb.RegRequest{ServerAddress: localAddress, ServerAlias: "storage01"})
	CheckError(err)
	log.Printf("Response from naming server: %s", response.GetStatus())
}
