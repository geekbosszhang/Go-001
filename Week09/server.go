package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listen: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening...")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn)  {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading ", err.Error())
	}
	conn.Write([]byte(fmt.Sprintf("Received %d charactors.", reqLen)))
	conn.Close()
}


