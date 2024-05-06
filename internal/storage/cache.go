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
	UpdateInWork([]service.OrderFull)
	NewOrder(order service.OrderFull)
}

type Cookie struct {
	Value  string
	Expire int64
}

var cache ICache

func CacheInit() error {
	orders, err := getInWorkOrders()
	if err != nil {
		return err
	}
	cache = heapcache.NewHeap(orders)
	return nil
}

func CacheUpdateInWork() error {
	data, err := getInWorkOrders()
	if err != nil {
		return err
	}
	cache.UpdateInWork(data)
	return nil
}

func CacheGetInWork() []service.OrderFull {
	return cache.GetInWork()
}

func CacheNewOrder(order service.OrderFull) {
	cache.NewOrder(order)
}
