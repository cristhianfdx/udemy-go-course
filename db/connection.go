package db

import (
	"context"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbURI = os.Getenv("CONNECTION_STRING")

// Mongo Variable connection
var MongoCN = ConnectDb()
var clientOptions = options.Client().ApplyURI(dbURI)

// Connection Function
func ConnectDb() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Success Connection")
	return client
}

// Check Database connection
func CheckConnection() bool {
	return MongoCN != nil
}
