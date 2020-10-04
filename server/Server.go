package main

import (
	"../pb"
	"bufio"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handle(conn net.Conn) {
	for {
		buf, err := bufio.NewReader(conn).ReadString('\n') // receive
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(buf)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		fmt.Print("-> ", buf)

		conn.Write([]byte("success")) //send
	}
}

func main() {
	//arguments := os.Args
	//port := ":" + arguments[1]
	port := ":5678"
	listener, err := net.Listen("tcp", port)
	CheckError(err)
	println("Listening on " + port)

	//aliveServer := NewKeepAliveServer()
	aliveServer := KeepAliveServerImpl{}
	grpcServer := grpc.NewServer()
	pb.RegisterKeepAliveServer(grpcServer, &aliveServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Errorf("error serving gRPC server %s", err)
		os.Exit(1)
	}
}
