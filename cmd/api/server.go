package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpserver "http101/cmd/api/http-server"
	routersetup "http101/cmd/api/router-setup"
	"http101/internal/application/middleware"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

// Run initializes the server and listens on the specified port.
func (s *Server) Run() {
	router := httpserver.NewRouter()
	addr := fmt.Sprintf("localhost%s", s.port)

	// Register endpoints.
	routersetup.New(router).RegisterTaskEndpoints()

	// Register middlewares.
	wrappedServer := middleware.LoggerMiddleware(router.Server)

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

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown...", sig)

	if err := httpServer.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	log.Println("Server stopped")
}
