package main

import (
	"net"
	"strings"
)

type Route struct {
	Path    string
	Method  HTTPMethod
	Handler func(conn net.Conn, params map[string]string)
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{routes: []Route{}}
}

func (r *Router) AddRoute(method, path string, handler func(conn net.Conn, params map[string]string)) {
	r.routes = append(r.routes, Route{Path: path, Method: HTTPMethod(method), Handler: handler})
}

func (r *Router) MatchRoute(method, path string) (handler func(conn net.Conn, params map[string]string), params map[string]string) {
	for _, route := range r.routes {
		if route.Method == HTTPMethod(method) {
			params, ok := matchPath(route.Path, path)

			if ok {
				return route.Handler, params
			}
		}
	}

	return nil, nil
}

func matchPath(routePath, requestPath string) (map[string]string, bool) {
	routeParts := strings.Split(routePath, "/")
	requestParts := strings.Split(requestPath, "/")

	if len(routeParts) != len(requestParts) {
		return nil, false
	}

	params := make(map[string]string)

	for i := range routeParts {
		if strings.HasPrefix(routeParts[i], "{") && strings.HasSuffix(routeParts[i], "}") {
			paramName := routeParts[i][1 : len(routeParts[i])-1]
			params[paramName] = requestParts[i]
		} else if routeParts[i] != requestParts[i] {
			return nil, false
		}
	}

	return params, true
}
