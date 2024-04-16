package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chemistry struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	SellValue  int                `json:"sellValue" bson:"sellValue"`
	ProbeValue int                `json:"probeValue" bson:"probeValue"`
	ProbeCount int                `json:"probeCount" bson:"probeCount"`
}

func (c *Chemistry) NewID() {
	c.ID = primitive.NewObjectID()
}

func (c *Chemistry) GetID() primitive.ObjectID {
	return c.ID
}

func (c *Chemistry) GetName() string {
	return c.Name
}
