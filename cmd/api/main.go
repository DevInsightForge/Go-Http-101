package main

import (
	"log"

	"http101/internal/application/config"
	"http101/internal/infrastructure/database"
	"http101/internal/webapi"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error initializing configuration variables: %v", err)
	}

	// Initialize MongoDB
	if err := database.InitMongoDB(config.MongoURI, config.MongoDbName); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Initialize API server
	server := webapi.NewServer(config.ServerAddr, config.ServerPort)
	server.Run()
}
