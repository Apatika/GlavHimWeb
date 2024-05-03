package heapcache

import (
	"sync"
)

type Cache struct {
	auth   map[string]string
	inWork interface{}
	mut    sync.RWMutex
}

func NewHeap(iwo interface{}) *Cache {
	return &Cache{
		auth:   make(map[string]string, 10),
		inWork: iwo,
	}
}

func (c *Cache) GetInWork() interface{} {
	return c.inWork
}

func (c *Cache) UpdateInWork(data interface{}) {
	c.mut.Lock()
	c.inWork = data
	c.mut.Unlock()
}
