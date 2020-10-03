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
	serverAddr := arguments[1] // "host:port" as a string
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr) // used by clients to establish connection
	checkError(err)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text+"\n") // send

		message, _ := bufio.NewReader(conn).ReadString('\n') // receive
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

	// receive a message
	//var buf [255]byte
	//_, err = conn.Read(buf[0:])
	//checkError(err)
	//fmt.Println(buf)
}
