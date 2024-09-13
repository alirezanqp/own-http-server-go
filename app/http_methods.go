package main

// HTTPMethod represents standard HTTP methods
type HTTPMethod string

// HTTP Methods as constants
const (
	HttpMethodGET     = "GET"
	HttpMethodPOST    = "POST"
	HttpMethodPUT     = "PUT"
	HttpMethodDELETE  = "DELETE"
	HttpMethodPATCH   = "PATCH"
	HttpMethodOPTIONS = "OPTIONS"
	HttpMethodHEAD    = "HEAD"
)
