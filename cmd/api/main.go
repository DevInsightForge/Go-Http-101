package main

import (
	"log"

	"http101/internal/infrastructure/database"
	"http101/internal/webapi"
)

func main() {
	// Initialize MongoDB
	if err := database.InitMongoDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Initialize API server
	server := webapi.NewServer()
	server.Run()
}
