package heapcache

import (
	"glavhim-app/internal/config"
	"sync"
)

// TODO: переделать в нормальный кэш
type Cache struct {
	inWork map[string]interface{} //костыль, чтобы не подгружать заказы в работе каждый раз из базы данных
	mut    sync.RWMutex
}

func New() *Cache {
	return &Cache{
		inWork: make(map[string]interface{}),
	}
}

func (c *Cache) Get(field string) map[string]interface{} {
	switch field {
	case config.Cfg.DB.Coll.Orders:
		return c.inWork
	default:
		return nil
	}
}

func (c *Cache) Update(field, key string, data interface{}) {
	c.mut.Lock()
	switch field {
	case config.Cfg.DB.Coll.Orders:
		c.inWork[key] = data
	}
	c.mut.Unlock()
}

func (c *Cache) Delete(field, key string) {
	c.mut.Lock()
	switch field {
	case config.Cfg.DB.Coll.Orders:
		delete(c.inWork, key)
	}
	c.mut.Unlock()
}
