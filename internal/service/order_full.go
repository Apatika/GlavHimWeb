package service

import (
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
)

type OrderFull struct {
	Customer Customer `json:"customer"`
	Order    Order    `json:"order"`
}

func NewOrderFull() OrderFull {
	return OrderFull{}
}

func (o *OrderFull) GetAll() interface{} {
	return storage.Cache.Get(config.Cfg.DB.Coll.Orders)
}
