package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"http101/internal/application/endpoint"
)

type Server struct {
	port string
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Before %s", r.URL.String())
		next.ServeHTTP(w, r)
		elapsedTime := time.Since(startTime)
		log.Printf("After %s - elapsed: %s", r.URL.String(), elapsedTime)
	})
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

// Run initializes the server and listens on the specified port.
func (s *Server) Run() {
	server := http.NewServeMux()
	addr := fmt.Sprintf("localhost%s", s.port)

	// Register endpoints.
	endpoint.RegisterTaskEndpoints(server)

	// Wrap server with logging middleware.
	wrappedServer := loggingMiddleware(server)

	// Setup server with options.
	httpServer := &http.Server{
		Addr:    addr,
		Handler: wrappedServer,
	}

	// Running server in a goroutine to allow graceful shutdown.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", addr, err)
		}
	}()

	log.Printf("Server is running at http://localhost%s\n", s.port)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	log.Println("Shutting down server...")

	if err := httpServer.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	log.Println("Server stopped")
}
