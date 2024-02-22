package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbConn *mongo.Database

type MongoDBConfig struct {
	ConnectionURI string
	DatabaseName  string
}

func InitMongoDB(config MongoDBConfig) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.ConnectionURI))
	if err != nil {
		return err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}

	dbConn = client.Database(config.DatabaseName)

	log.Println("Connected to MongoDB")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	if dbConn == nil {
		log.Fatal("Database connection is not initialized")
	}
	return dbConn.Collection(collectionName)
}
