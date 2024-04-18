package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IService interface {
	NewID()
	GetID() primitive.ObjectID
	GetName() string
}

type Pvd struct {
	Weight float32 `json:"weight" bson:"weight"`
	Count  int     `json:"count" bson:"count"`
}

type Date struct {
	Day   int `json:"day" bson:"day"`
	Month int `json:"month" bson:"month"`
	Year  int `json:"year" bson:"year"`
}

type Contact struct {
	Name string `json:"name" bson:"name"`
	Tel  string `json:"tel" bson:"tel"`
}

type Adress struct {
	City     string `json:"city" bson:"city"`
	Adress   string `json:"adress" bson:"adress"`
	Terminal string `json:"terminal" bson:"terminal"`
}
