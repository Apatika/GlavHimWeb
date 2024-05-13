package heapcache

import (
	"glavhim-app/internal/config"
	"sync"
)

type Cache struct {
	auth      map[string]string
	inWork    map[string]interface{}
	chemistry map[string]interface{}
	cargos    map[string]interface{}
	users     map[string]interface{}
	mut       sync.RWMutex
}

func New() *Cache {
	return &Cache{
		auth:      make(map[string]string),
		inWork:    make(map[string]interface{}),
		chemistry: make(map[string]interface{}),
		cargos:    make(map[string]interface{}),
		users:     make(map[string]interface{}),
	}
}

func (c *Cache) Get(field string) map[string]interface{} {
	switch field {
	case config.Cfg.DB.Coll.Orders:
		return c.inWork
	case config.Cfg.DB.Coll.Cargos:
		return c.cargos
	case config.Cfg.DB.Coll.Chemistry:
		return c.chemistry
	case config.Cfg.DB.Coll.Users:
		return c.users
	default:
		return nil
	}
}

func (c *Cache) Update(field, key string, data interface{}) {
	c.mut.Lock()
	switch field {
	case config.Cfg.DB.Coll.Orders:
		c.inWork[key] = data
	case config.Cfg.DB.Coll.Cargos:
		c.cargos[key] = data
	case config.Cfg.DB.Coll.Chemistry:
		c.chemistry[key] = data
	case config.Cfg.DB.Coll.Users:
		c.users[key] = data
	}
	c.mut.Unlock()
}

func (c *Cache) Delete(field, key string) {
	c.mut.Lock()
	switch field {
	case config.Cfg.DB.Coll.Orders:
		delete(c.inWork, key)
	case config.Cfg.DB.Coll.Cargos:
		delete(c.cargos, key)
	case config.Cfg.DB.Coll.Chemistry:
		delete(c.chemistry, key)
	case config.Cfg.DB.Coll.Users:
		delete(c.users, key)
	}
	c.mut.Unlock()
}
