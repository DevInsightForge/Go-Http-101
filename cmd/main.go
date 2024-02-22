package main

import (
	"log"

	"http101/cmd/api"
	"http101/internal/application/config"
	"http101/internal/application/database"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize MongoDB
	if err := database.InitMongoDB(config.MongoURI, config.MongoDbName); err != nil {
		log.Fatalf("Error initializing MongoDB: %v", err)
	}

	// Initialize API server
	server := api.NewServer(config.ServerAddr, config.ServerPort)
	server.Run()
}
