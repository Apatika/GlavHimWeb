package storage

import (
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/heapcache"
)

const (
	IWO = "inWorkOrders"
	COO = "cookie"
)

type ICache interface {
	GetInWork() []service.OrderFull
	NewOrder(service.OrderFull)
	UpdateOrder(service.OrderFull) error
	UpdateOrderStatus(service.OrderStatusChanger) error
}

type Cookie struct {
	Value  string
	Expire int64
}

var cache ICache

func CacheInit() error {
	db := DB()
	orders, err := db.GetInWorkOrders()
	if err != nil {
		return err
	}
	cache = heapcache.NewHeap(orders)
	return nil
}

func CacheGetInWork() []service.OrderFull {
	return cache.GetInWork()
}

func CacheNewOrder(order service.OrderFull) {
	cache.NewOrder(order)
}

func CacheUpdateOrder(order service.OrderFull) error {
	if err := cache.UpdateOrder(order); err != nil {
		return err
	}
	return nil
}

func CacheUpdateOrderStatus(status service.OrderStatusChanger) error {
	if err := cache.UpdateOrderStatus(status); err != nil {
		return err
	}
	return nil
}
