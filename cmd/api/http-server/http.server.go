package httpserver

import (
	"fmt"
	"net/http"
)

type Router struct {
	Server *http.ServeMux
	prefix string
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		Server: http.NewServeMux(),
		prefix: "/api",
	}
}

func (r *Router) Handle(method string, path string, handler http.HandlerFunc) {
	newPath := fmt.Sprintf("%s %s", method, r.prefix+path)
	r.Server.HandleFunc(newPath, handler)
}

func (r *Router) Group(prefix string) *Router {
	return &Router{
		Server: r.Server,
		prefix: r.prefix + prefix,
	}
}
