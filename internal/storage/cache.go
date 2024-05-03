package storage

import (
	"glavhim-app/internal/storage/heapcache"
)

const (
	IWO = "inWorkOrders"
	COO = "cookie"
)

type ICache interface {
	GetInWork() interface{}
	UpdateInWork(interface{})
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

func CacheGetInWork() interface{} {
	return cache.GetInWork()
}
