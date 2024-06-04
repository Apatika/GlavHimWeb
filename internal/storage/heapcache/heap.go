package heapcache

import (
	"sync"
)

type Cache struct {
	Data map[string]interface{}
	mut  sync.RWMutex
}

func New() *Cache {
	return &Cache{
		Data: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) interface{} {
	return c.Data[key]
}

func (c *Cache) Update(key string, data interface{}) {
	c.mut.Lock()
	c.Data[key] = data
	c.mut.Unlock()
}

func (c *Cache) Delete(key string) {
	c.mut.Lock()
	delete(c.Data, key)
	c.mut.Unlock()
}
