package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Money struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Day     int                `json:"day" bson:"day"`
	Month   int                `json:"month" bson:"month"`
	Year    int                `json:"year" bson:"year"`
	Price   int                `json:"price" bson:"price"`
	Comment string             `json:"comment" bson:"comment"`
}
