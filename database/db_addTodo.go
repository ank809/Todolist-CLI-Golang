package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ank809/Todolist-CLI-Golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTodo(title string, description string) (string, error) {

	id := primitive.NewObjectID()
	var todo models.Todo = models.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}

	collection_name := "todos"
	collection := OpenCollection(MongoClient, collection_name)
	_, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		fmt.Println("Error in inserting todo")
		return "", err
	}
	return "Todo Successfully Added", err

}
