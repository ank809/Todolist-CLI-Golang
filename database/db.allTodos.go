package database

import (
	"context"
	"fmt"

	"github.com/ank809/Todolist-CLI-Golang/models"
	"gopkg.in/mgo.v2/bson"
)

func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo

	collection_name := "todos"

	collection := OpenCollection(MongoClient, collection_name)
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return todos, err
	}
	if err = cursor.All(context.Background(), &todos); err != nil {
		fmt.Println(err)
		return todos, err
	}
	return todos, nil
}
