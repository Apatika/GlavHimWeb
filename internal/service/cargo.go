package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cargo struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	URI        string             `json:"uri" bson:"uri"`
	MainTel    string             `json:"mainTel" bson:"mainTel"`
	ManagerTel string             `json:"managerTel" bson:"managerTel"`
}

func (c *Cargo) NewID() {
	c.ID = primitive.NewObjectID()
}

func (c *Cargo) GetID() primitive.ObjectID {
	return c.ID
}

func (c *Cargo) GetName() string {
	return c.Name
}
