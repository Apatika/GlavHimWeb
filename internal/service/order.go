package service

import (
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
	"log"
)

type Order struct {
	ID           string               `json:"id" bson:"_id"`
	CreationDate Date                 `json:"creationDate" bson:"creation_date"`
	LastDate     string               `json:"lastDate" bson:"last_date"`
	Cargo        string               `json:"cargo" bson:"cargo"`
	ToAdress     bool                 `json:"toadress" bson:"toadress"`
	Adress       Adress               `json:"adress" bson:"adress"`
	ClientID     string               `json:"clientId" bson:"client_id"`
	Invoice      []string             `json:"invoice" bson:"invoice"`
	Comment      string               `json:"comment" bson:"comment"`
	Probes       map[string]Chemistry `json:"probes" bson:"probes"`
	Payment      bool                 `json:"payment" bson:"payment"`
	ShipmentDate Date                 `json:"shipmentDate" bson:"shipment_date"`
	Status       string               `json:"status" bson:"status"`
}

func (o *Order) Push() error {
	db := storage.DB()
	if err := db.Add(config.Cfg.DB.Coll.Orders, o); err != nil {
		return fmt.Errorf("write order to db failed(%v)", err.Error())
	}
	log.Printf("create order (id: %v)", o.ID)
	return nil
}

func (o *Order) Update() error {
	db := storage.DB()
	if err := db.Update(config.Cfg.DB.Coll.Orders, o, o.ID); err != nil {
		return fmt.Errorf("update order failed(%v)", err.Error())
	}
	log.Printf("update order (id: %v)", o.ID)
	return nil
}

func (o *Order) Delete() error {
	db := storage.DB()
	if err := db.Delete(config.Cfg.DB.Coll.Orders, o.ID); err != nil {
		return fmt.Errorf("delete order from db failed (%v)", err.Error())
	}
	storage.Cache.Delete(config.Cfg.DB.Coll.Orders, o.ID)
	log.Printf("delete order (id: %v)", o.ID)
	return nil
}
