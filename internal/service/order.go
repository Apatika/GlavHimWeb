package service

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
	"log"
	"time"
)

type Order struct {
	ID           string               `json:"id" bson:"_id"`
	CreationDate string               `json:"creationDate" bson:"creation_date"`
	LastDate     string               `json:"lastDate" bson:"last_date"`
	Cargo        string               `json:"cargo" bson:"cargo"`
	ToAdress     bool                 `json:"toadress" bson:"toadress"`
	Adress       Adress               `json:"adress" bson:"adress"`
	CustomerID   string               `json:"customerId" bson:"customer_id"`
	Invoice      []string             `json:"invoice" bson:"invoice"`
	Comment      string               `json:"comment" bson:"comment"`
	Probes       map[string]Chemistry `json:"probes" bson:"probes"`
	Payment      bool                 `json:"payment" bson:"payment"`
	ShipmentDate string               `json:"shipmentDate" bson:"shipment_date"`
	Status       string               `json:"status" bson:"status"`
	RouteNum     string               `json:"routeNum" bson:"routeNum"`
	Customer     *Customer            `json:"customer" bson:"customer,omitempty"`
	Content      string               `json:"content" bson:"content"`
}

func InWorkToCache() error {
	var orders []Order
	if err := storage.DB.InWorkOrders(&orders); err != nil {
		return err
	}
	storage.Cache.Update(config.Cfg.DB.Coll.Orders, orders)
	return nil
}

func GetInWorkOrders() interface{} {
	return storage.Cache.Get(config.Cfg.DB.Coll.Orders)
}

func (o *Order) Push(raw []byte) error {
	json.Unmarshal(raw, o)
	o.CreationDate = time.Now().Format(time.DateTime)
	o.ID = storage.DB.GetNewID()
	if o.Customer.ID == "" {
		id, err := o.Customer.Check()
		if err != nil || id == "" {
			o.Customer.ID = storage.DB.GetNewID()
			if err := o.Customer.Push(); err != nil {
				return fmt.Errorf("push customer failed(%v)", err.Error())
			}
		} else {
			o.Customer.ID = id
			if err := o.Customer.Update(); err != nil {
				return fmt.Errorf("update customer failed(%v)", err.Error())
			}
		}
	} else {
		if err := o.Customer.Update(); err != nil {
			return fmt.Errorf("update customer failed(%v)", err.Error())
		}
	}
	o.CustomerID = o.Customer.ID
	o.Customer = nil
	if err := storage.DB.Add(config.Cfg.DB.Coll.Orders, o); err != nil {
		return fmt.Errorf("write order to db failed(%v)", err.Error())
	}
	log.Printf("create order (id: %v)", o.ID)
	InWorkToCache()
	return nil
}

func (o *Order) Update(raw []byte) error {
	json.Unmarshal(raw, o)
	if o.ID == "" {
		return fmt.Errorf("zero order id")
	}
	if o.Status == StatusShipped {
		o.ShipmentDate = time.Now().Format(time.DateTime)
	} else {
		o.ShipmentDate = ""
	}
	if err := o.Customer.Update(); err != nil {
		return err
	}
	o.Customer = nil
	if err := storage.DB.Update(config.Cfg.DB.Coll.Orders, o, o.ID); err != nil {
		return fmt.Errorf("update order failed(%v)", err.Error())
	}
	log.Printf("update order (id: %v)", o.ID)
	InWorkToCache()
	return nil
}

func (o *Order) Delete() error {
	if err := storage.DB.Delete(config.Cfg.DB.Coll.Orders, o.ID); err != nil {
		return fmt.Errorf("delete order from db failed (%v)", err.Error())
	}
	log.Printf("delete order (id: %v)", o.ID)
	return nil
}
