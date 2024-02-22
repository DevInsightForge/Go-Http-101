package httpserver

import (
	"net/http"
	"strings"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Router struct {
	Mux         *http.ServeMux
	middlewares []Middleware
	prefix      string
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		Mux:    http.NewServeMux(),
		prefix: "",
	}
}

func (r *Router) Handle(path string, handler http.HandlerFunc) {
	newPath := strings.Split(path, " ")[0] + " " + r.prefix + strings.Split(path, " ")[1]
	r.Mux.HandleFunc(newPath, r.applyMiddleware(handler))
}

func (r *Router) Group(prefix string) *Router {
	return &Router{
		Mux:    r.Mux,
		prefix: r.prefix + prefix,
	}
}

func (r *Router) Use(middleware ...Middleware) {
	r.middlewares = append(r.middlewares, middleware...)
}

func (r *Router) applyMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	finalHandler := handler
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		finalHandler = r.middlewares[i](finalHandler)
	}
	return finalHandler
}
