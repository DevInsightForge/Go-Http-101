package main

import (
	"fmt"
	"log"
	"net/http"

	"http101/internal/application/config"
	"http101/internal/application/database"
	"http101/internal/application/endpoint"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize MongoDB
	mongoConfig := database.MongoDBConfig{
		ConnectionURI: config.MongoURI,
		DatabaseName:  config.MongoDbName,
	}

	if err := database.InitMongoDB(mongoConfig); err != nil {
		log.Fatalf("Error initializing MongoDB: %v", err)
	}

	server := http.NewServeMux()
	endpoint.RegisterTaskEndpoints(server)

	fmt.Printf("Server is running at http://localhost%s\n", config.ServerPort)

	err = http.ListenAndServe(fmt.Sprintf("localhost%s", config.ServerPort), server)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
