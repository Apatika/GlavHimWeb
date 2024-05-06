package heapcache

import (
	"glavhim-app/internal/service"
	"sync"
)

type Cache struct {
	auth   map[string]string
	inWork []service.OrderFull
	mut    sync.RWMutex
}

func NewHeap(iwo []service.OrderFull) *Cache {
	return &Cache{
		auth:   make(map[string]string, 10),
		inWork: iwo,
	}
}

func (c *Cache) GetInWork() []service.OrderFull {
	return c.inWork
}

func (c *Cache) UpdateInWork(data []service.OrderFull) {
	c.mut.Lock()
	c.inWork = data
	c.mut.Unlock()
}

func (c *Cache) NewOrder(order service.OrderFull) {

}
