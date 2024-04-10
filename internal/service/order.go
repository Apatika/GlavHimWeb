package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pvd struct {
	Weight float32 `json:"weight" bson:"weight"`
	Count  int     `json:"count" bson:"count"`
}

type Date struct {
	Day   int `json:"day" bson:"day"`
	Month int `json:"month" bson:"month"`
	Year  int `json:"year" bson:"year"`
}

type Order struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CreationDate Date               `json:"creationDate" bson:"creation_date"`
	LastDate     string             `json:"lastDate" bson:"last_date"`
	Cargo        string             `json:"cargo" bson:"cargo"`
	Adress       Adress             `json:"adress" bson:"adress"`
	ClientID     int                `json:"clientId" bson:"client_id"`
	Invoice      []int              `json:"invoice" bson:"invoice"`
	Comment      string             `json:"comment" bson:"comment"`
	Probes       []Chemistry        `json:"probes" bson:"probes"`
	Pvd          []Pvd              `json:"pvd" bson:"pvd"`
	Payment      bool               `json:"payment" bson:"payment"`
	ShipmentDate Date               `json:"shipmentDate" bson:"shipment_date"`
	Status       int                `json:"status" bson:"status"`
}
