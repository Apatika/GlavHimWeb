package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Tel   string             `json:"tel" bson:"tel"`
	Email string             `json:"email" bson:"email"`
}
