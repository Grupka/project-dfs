package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	arguments := os.Args
	port := ":" + arguments[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// accept an incoming connection
	conn, err := listener.Accept()
	checkError(err)

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
		fmt.Print("-> ", string(buf))

		conn.Write([]byte("success")) //send
	}
	listener.Close()
}
