package storage

import (
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/heapcache"
)

type ICache interface {
	Get(string) map[string]interface{}
	Update(string, string, interface{})
	Delete(string, string)
}

var Cache ICache

func CacheInit() {
	Cache = heapcache.New()
	load()
}

func load() error {
	db := DB()
	orders, err := db.GetInWorkOrders()
	if err != nil {
		return fmt.Errorf("cache load error")
	}
	for _, v := range orders {
		Cache.Update(config.Cfg.DB.Coll.Orders, v.Order.ID, v)
	}
	var cargos []service.Cargo
	if err := db.GetAll(config.Cfg.DB.Coll.Cargos, &cargos); err != nil {
		return err
	}
	for _, v := range cargos {
		Cache.Update(config.Cfg.DB.Coll.Cargos, v.ID, v)
	}
	var chemistry []service.Chemistry
	if err := db.GetAll(config.Cfg.DB.Coll.Chemistry, &chemistry); err != nil {
		return err
	}
	for _, v := range chemistry {
		Cache.Update(config.Cfg.DB.Coll.Chemistry, v.ID, v)
	}
	var users []service.User
	if err := db.GetAll(config.Cfg.DB.Coll.Users, &users); err != nil {
		return err
	}
	for _, v := range users {
		Cache.Update(config.Cfg.DB.Coll.Users, v.ID, v)
	}
	return nil
}
