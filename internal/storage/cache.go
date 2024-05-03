package storage

import (
	"sync"
)

const (
	IWO = "inWorkOrders"
	COO = "cookie"
)

type Cookie struct {
	Value  string
	Expire int64
}

var Cache = make(map[string]interface{})
var cacheMut sync.RWMutex

func CacheInit() error {
	Cache[COO] = make([]Cookie, 0, 5)
	if err := CacheUpdateInWork(); err != nil {
		return err
	}
	return nil
}

func CacheUpdateInWork() error {
	inWork, err := getInWorkOrders()
	if err != nil {
		return err
	}
	cacheMut.Lock()
	Cache[IWO] = inWork
	cacheMut.Unlock()
	return nil
}

func InWork() interface{} {
	data := Cache[IWO]
	return data
}
