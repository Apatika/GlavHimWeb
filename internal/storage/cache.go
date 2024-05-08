package storage

import (
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/heapcache"
)

type ICache interface {
	GetInWork() []service.OrderFull
	NewOrder(service.OrderFull)
	UpdateOrder(service.OrderFull) error
	UpdateOrderStatus(service.OrderStatusChanger) error
}

func Cache() (ICache, error) {
	db := DB()
	orders, err := db.GetInWorkOrders()
	if err != nil {
		return nil, err
	}
	return heapcache.New(orders), nil
}
