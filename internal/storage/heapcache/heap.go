package heapcache

import (
	"fmt"
	"glavhim-app/internal/service"
	"sync"
)

type Cache struct {
	auth     map[string]string
	inWork   []service.OrderFull
	chems    map[string]service.Chemistry
	cargos   map[string]service.Cargo
	managers map[string]service.User
	mut      sync.RWMutex
}

func New(iwo []service.OrderFull, chems map[string]service.Chemistry, cargos map[string]service.Cargo, managers map[string]service.User) *Cache {
	return &Cache{
		auth:     make(map[string]string, 10),
		inWork:   iwo,
		chems:    chems,
		cargos:   cargos,
		managers: managers,
	}
}

func (c *Cache) GetInWork() []service.OrderFull {
	return c.inWork
}

func (c *Cache) NewOrder(order service.OrderFull) {
	c.mut.Lock()
	c.inWork = append(c.inWork, order)
	c.mut.Unlock()
}

func (c *Cache) UpdateOrder(order service.OrderFull) error {
	c.mut.Lock()
	is := false
	for i, v := range c.inWork {
		if order.Order.ID == v.Order.ID {
			is = true
			c.inWork[i] = order
		}
	}
	c.mut.Unlock()
	if !is {
		return fmt.Errorf("not found order (ID: %v) in cache", order.Order.ID)
	}
	return nil
}

func (c *Cache) UpdateOrderStatus(status service.OrderStatusChanger) error {
	c.mut.Lock()
	is := false
	for i, v := range c.inWork {
		if status.ID == v.Order.ID {
			is = true
			c.inWork[i].Order.Status = status.Status
			c.inWork[i].Order.ShipmentDate = status.ShipmentDate
		}
	}
	c.mut.Unlock()
	if !is {
		return fmt.Errorf("not found order (ID: %v) in cache", status.ID)
	}
	return nil
}
