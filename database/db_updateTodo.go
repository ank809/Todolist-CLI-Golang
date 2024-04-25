package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ank809/Todolist-CLI-Golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func UpdateTodo(Id string, title string, description string) (string, error) {
	objectid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		fmt.Println(err)
		return "Error", err
	}
	var todo models.Todo = models.Todo{
		ID:          objectid,
		Title:       title,
		Description: description,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
	updatedtodo := bson.M{"$set": todo}
	filter := bson.M{"_id": objectid}
	collection_name := "todos"
	collection := OpenCollection(MongoClient, collection_name)
	res := collection.FindOneAndUpdate(context.Background(), filter, updatedtodo)
	if res.Err() != nil {
		fmt.Println(res.Err())
		return "", res.Err()
	}
	return "Todo Updated successfully", nil
}
