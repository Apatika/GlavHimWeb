package service

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
	"log"
	"time"
)

type OrderFull struct {
	Client Client `json:"client"`
	Order  Order  `json:"order"`
}

func NewOrderFull() OrderFull {
	return OrderFull{}
}

func (o *OrderFull) GetAll() map[string]interface{} {
	return storage.Cache.Get(config.Cfg.DB.Coll.Orders)
}

func (o *OrderFull) Push(raw []byte) error {
	db := storage.DB()
	json.Unmarshal(raw, o)
	var month time.Month
	o.Order.CreationDate.Year, month, o.Order.CreationDate.Day = time.Now().Date()
	o.Order.CreationDate.Month = month.String()
	o.Order.ID = db.GetNewID()
	if o.Client.ID == "" {
		id, err := o.Client.Check()
		if err != nil || id == "" {
			o.Client.ID = db.GetNewID()
			if err := o.Client.Push(); err != nil {
				return fmt.Errorf("push client failed(%v)", err.Error())
			}
		} else {
			o.Client.ID = id
			if err := o.Client.Update(); err != nil {
				return fmt.Errorf("update client failed(%v)", err.Error())
			}
		}
	} else {
		if err := o.Client.Update(); err != nil {
			return fmt.Errorf("update client failed(%v)", err.Error())
		}
	}
	o.Order.ClientID = o.Client.ID
	if err := o.Order.Push(); err != nil {
		return fmt.Errorf("push order failed(%v)", err.Error())
	}
	storage.Cache.Update(config.Cfg.DB.Coll.Orders, o.Order.ID, o)
	return nil
}

func (o *OrderFull) Update(raw []byte) error {
	json.Unmarshal(raw, o)
	if o.Order.ID == "" {
		return fmt.Errorf("zero order id")
	}
	if o.Order.Status == StatusShipped {
		var month time.Month
		o.Order.ShipmentDate.Year, month, o.Order.ShipmentDate.Day = time.Now().Date()
		o.Order.ShipmentDate.Month = month.String()
	}
	if err := o.Client.Update(); err != nil {
		return err
	}
	if err := o.Order.Update(); err != nil {
		return err
	}
	if o.Order.Status == StatusShipped {
		storage.Cache.Delete(config.Cfg.DB.Coll.Orders, o.Order.ID)
		log.Printf("update status ID: %v, status: %v", o.Order.ID, o.Order.Status)
	} else {
		storage.Cache.Update(config.Cfg.DB.Coll.Orders, o.Order.ID, o)
		log.Printf("update order ID: %v", o.Order.ID)
	}
	return nil
}
