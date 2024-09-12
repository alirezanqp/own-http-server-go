package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 4221")

	for {
		// Block until we receive an incoming connection
		req, err := listener.Accept()

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client request
		handleRequest(req)
	}
}

func handleRequest(req net.Conn) {
	// Ensure we close the connection after we're done
	defer req.Close()

	checkRequestUrlExists(req)
}

func checkRequestUrlExists(req net.Conn) {
	reqData := make([]byte, 1024)

	_, err := req.Read(reqData)

	if err != nil {
		return
	}

	log.Println(string(reqData))

	if !strings.HasPrefix(string(reqData), "GET / HTTP/1.1") {
		req.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	} else {
		req.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	}
}
