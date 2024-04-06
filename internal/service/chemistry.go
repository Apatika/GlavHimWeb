package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chemistry struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string
	SellValue  int     `json:"sellValue" bson:"sell_value"`
	ProbeValue float32 `json:"probeValue" bson:"probe_value"`
	ProbeCount int     `json:"probeCount" bson:"probe_count"`
}
