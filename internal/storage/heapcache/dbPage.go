package heapcache

import (
	"glavhim-app/internal/service"
)

func (c *Cache) Managers() []service.User {
	var arr []service.User
	for _, v := range c.managers {
		arr = append(arr, v)
	}
	return arr
}

func (c *Cache) Cargos() []service.Cargo {
	var arr []service.Cargo
	for _, v := range c.cargos {
		arr = append(arr, v)
	}
	return arr
}

func (c *Cache) Chemistry() []service.Chemistry {
	var arr []service.Chemistry
	for _, v := range c.chems {
		arr = append(arr, v)
	}
	return arr
}

func (c *Cache) AddToManagers(data service.User) {
	c.mut.Lock()
	c.managers[data.ID] = data
	c.mut.Unlock()
}

func (c *Cache) AddToCargos(data service.Cargo) {
	c.mut.Lock()
	c.cargos[data.ID] = data
	c.mut.Unlock()
}

func (c *Cache) AddToChemistry(data service.Chemistry) {
	c.mut.Lock()
	c.chems[data.ID] = data
	c.mut.Unlock()
}

func (c *Cache) DeleteManagers(data service.User) {
	c.mut.Lock()
	delete(c.managers, data.ID)
	c.mut.Unlock()
}

func (c *Cache) DeleteCargos(data service.Cargo) {
	c.mut.Lock()
	delete(c.cargos, data.ID)
	c.mut.Unlock()
}

func (c *Cache) DeleteChemistry(data service.Chemistry) {
	c.mut.Lock()
	delete(c.chems, data.ID)
	c.mut.Unlock()
}
