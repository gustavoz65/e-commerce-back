package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to ping database:\t", err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	return client

}

var Client *mongo.Client = InitDB()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collections *mongo.Collection = client.Database("E-Commerce").Collection(collectionName)
	return collections
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("E-Commerce").Collection(collectionName)
	return productCollection
}
