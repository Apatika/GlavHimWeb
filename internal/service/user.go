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

func (u *User) NewID() {
	u.ID = primitive.NewObjectID()
}

func (u *User) GetID() primitive.ObjectID {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}
