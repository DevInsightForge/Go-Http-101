package config

import (
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

// Config represents the server configuration
type Config struct {
	ServerAddr  string `json:"address"`
	ServerPort  string `json:"port"`
	MongoURI    string `json:"mongouri"`
	MongoDbName string `json:"mongodbname"`
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	addr := getenv("ADDRESS", "localhost")
	port := ":" + getenv("PORT", "4000")
	mongoURI := getenv("MONGO_URI", "mongodb://localhost:27017")
	databaseName := getenv("DATABASE_NAME", "go101")

	return &Config{
		ServerAddr:  addr,
		ServerPort:  port,
		MongoURI:    mongoURI,
		MongoDbName: databaseName,
	}, nil
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
