package main

import (
	"fmt"
	"net"
)

// Utility function to send an HTTP response
func sendResponse(conn net.Conn, body string, statusCode HTTPStatus) {
	code := 200
	if statusCode > 0 {
		code = int(statusCode)
	}

	var response string

	if code >= 100 && code < 400 {
		response = fmt.Sprintf("HTTP/1.1 %d OK\r\n", code)
	} else {
		response = fmt.Sprintf("HTTP/1.1 %d\r\n", code)
	}

	response += "Content-Type: text/plain\r\n"
	response += fmt.Sprintf("Content-Length: %d\r\n", len(body))
	response += "\r\n"
	response += body

	// Send the response
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}
