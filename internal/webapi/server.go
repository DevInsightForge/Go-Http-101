package webapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"http101/internal/webapi/endpoint"
)

type Server struct {
	addr string
	port string
}

func NewServer(addr string, port string) *Server {
	return &Server{
		addr: addr,
		port: port,
	}
}

func (s *Server) Run() {
	handler := chi.NewRouter()

	// global middleware registration
	handler.Use(middleware.Logger)

	// health check endpoint
	handler.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK!"))
	})

	// API routes registration
	handler.Route("/api", func(apiRoute chi.Router) {
		apiRoute.Mount("/tasks", endpoint.TaskEndpoint{}.Routes())
	})

	// Setup server with options.
	fullAddr := fmt.Sprintf("%s:%s", s.addr, s.port)
	httpServer := &http.Server{
		Addr:    fullAddr,
		Handler: handler,
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	log.Printf("Server is running at http://%s\n", httpServer.Addr)

	// Run the server
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	log.Println("Server stopped")
}
