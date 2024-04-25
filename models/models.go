package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	CreatedAt   primitive.DateTime `json:"createdAt"`
}
