package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CreationDate Date               `json:"creationDate" bson:"creation_date"`
	LastDate     string             `json:"lastDate" bson:"last_date"`
	Cargo        string             `json:"cargo" bson:"cargo"`
	ToAdress     bool               `json:"toadress" bson:"toadress"`
	Adress       Adress             `json:"adress" bson:"adress"`
	ClientID     primitive.ObjectID `json:"clientId" bson:"client_id"`
	Invoice      []string           `json:"invoice" bson:"invoice"`
	Comment      string             `json:"comment" bson:"comment"`
	Probes       []Chemistry        `json:"probes" bson:"probes"`
	Payment      bool               `json:"payment" bson:"payment"`
	ShipmentDate Date               `json:"shipmentDate" bson:"shipment_date"`
	Status       string             `json:"status" bson:"status"`
}

func (o *Order) NewID() {
	o.ID = primitive.NewObjectID()
}
