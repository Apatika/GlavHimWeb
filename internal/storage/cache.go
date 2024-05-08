package storage

import (
	"glavhim-app/internal/config"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/heapcache"
)

type ICache interface {
	GetInWork() []service.OrderFull
	NewOrder(service.OrderFull)
	UpdateOrder(service.OrderFull) error
	UpdateOrderStatus(service.OrderStatusChanger) error
	Managers() []service.User
	Cargos() []service.Cargo
	Chemistry() []service.Chemistry
}

func Cache() (ICache, error) {
	db := DB()
	orders, err := db.GetInWorkOrders()
	if err != nil {
		return nil, err
	}
	var chems []service.Chemistry
	if err := db.GetAll(config.Cfg.DB.Coll.Chems, &chems); err != nil {
		return nil, err
	}
	var cargos []service.Cargo
	if err := db.GetAll(config.Cfg.DB.Coll.Cargos, &cargos); err != nil {
		return nil, err
	}
	var users []service.User
	if err := db.GetAll(config.Cfg.DB.Coll.Users, &users); err != nil {
		return nil, err
	}
	return heapcache.New(orders, chems, cargos, users), nil
}
