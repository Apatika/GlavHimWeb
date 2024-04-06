package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pvd struct {
	Weight float32
	Count  int
}

type CreationDate struct {
	Day   int
	Month int
	Year  int
}

type ShipmentDate struct {
	Day   int
	Month int
	Year  int
}

type Order struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CreationDate CreationDate       `json:"creationDate" bson:"creation_date"`
	LastDate     string             `json:"lastDate" bson:"last_date"`
	Cargo        string             `json:"cargo" bson:"cargo"`
	ClientID     int                `json:"clientId" bson:"client_id"`
	Invoice      []int
	Comment      string
	Probes       []Chemistry
	Pvd          []Pvd
	Payment      bool
	ShipmentDate ShipmentDate `json:"shipmentDate" bson:"shipment_date"`
	Status       int
}
