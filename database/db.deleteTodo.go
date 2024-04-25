package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func Deletetodos(id string) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	filter := bson.M{"_id": objectId}
	collection_name := "todos"
	collection := OpenCollection(MongoClient, collection_name)
	res := collection.FindOneAndDelete(context.Background(), filter)
	if res.Err() != nil {
		return "", err
	}
	return "Todo deleted successfully", nil
}
