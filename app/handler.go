package main

import (
	"fmt"
	"net"
)

// Example Handlers
func homeHandler(conn net.Conn, params map[string]string) {
	sendResponse(conn, "Welcome to the home page!", StatusOK)
}

func aboutHandler(conn net.Conn, params map[string]string) {
	sendResponse(conn, "This is the about page!", StatusOK)
}

func userHandler(conn net.Conn, params map[string]string) {
	sendResponse(conn, fmt.Sprintf("User ID: %s", params["id"]), StatusOK)
}
