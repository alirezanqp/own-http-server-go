package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// Create a new router
	router := &Router{}

	// Define routes
	router.AddRoute(HttpMethodGET, "/", homeHandler)
	router.AddRoute(HttpMethodGET, "/about", aboutHandler)
	router.AddRoute(HttpMethodGET, "/users/{id}", userHandler)

	// Start TCP server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Listening on :8080")

	for {
		// Accept new connections
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, router)
	}
}

func handleConnection(conn net.Conn, router *Router) {
	defer conn.Close()

	// Read request from the connection
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	// Parse HTTP request line ("GET /path HTTP/1.1")
	parts := strings.Split(requestLine, " ")
	if len(parts) < 3 {
		fmt.Println("Invalid request line")
		return
	}

	method := parts[0]
	path := parts[1]

	fmt.Printf("Received request: %s %s\n", method, path)

	// Match the request to a route
	handler, params := router.MatchRoute(method, path)
	if handler != nil {
		handler(conn, params)
	} else {
		sendResponse(conn, "404 Not Found", StatusNotFound)
	}
}
