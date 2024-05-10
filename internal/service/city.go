package service

type City struct {
	Region string `json:"region" bson:"region"`
	City   string `json:"city" bson:"city"`
}
