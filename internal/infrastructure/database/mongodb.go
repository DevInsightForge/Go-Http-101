package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"http101/internal/application/config"
)

var dbConn *mongo.Database

type MongoDBConfig struct {
	ConnectionURI string
	DatabaseName  string
}

func InitMongoDB() error {
	cfg := config.GetConfig()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}

	dbConn = client.Database(cfg.MongoDbName)
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	if dbConn == nil {
		log.Fatal("Database connection is not initialized yet")
	}

	return dbConn.Collection(collectionName)
}
