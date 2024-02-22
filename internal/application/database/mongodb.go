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

func InitMongoDB(ConnectionURI string, DatabaseName  string) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(ConnectionURI))
	if err != nil {
		return err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}

	dbConn = client.Database(DatabaseName)

	log.Println("Connection to MongoDB was initialized successfully")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	if dbConn == nil {
		log.Fatal("Database connection is not initialized yet")
	}
	return dbConn.Collection(collectionName)
}
