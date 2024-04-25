package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client = DbClient()

func DbClient() *mongo.Client {
	url := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println(err)
	}
	return client
}

func OpenCollection(client *mongo.Client, collection_name string) *mongo.Collection {
	collection := client.Database("Todolist-Cli").Collection(collection_name)
	return collection
}
