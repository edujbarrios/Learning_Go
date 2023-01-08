package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	// Use a bufio.Reader to read data from the connection
	reader := bufio.NewReader(conn)

	// Read data from the connection until there is no more
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// Echo the data back to the client
		conn.Write([]byte(data))
	}
}

func main() {
	// Bind to port 5000 on all available interfaces
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	// Accept incoming connections and handle them in a new goroutine
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
