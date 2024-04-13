package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IService interface {
	NewID()
	GetID() primitive.ObjectID
	GetName() string
}
