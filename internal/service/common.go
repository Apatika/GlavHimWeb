package service

import (
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
)

func GetCities(reg string) ([]City, error) {
	var cities []City
	db := storage.DB()
	if err := db.GetCities(reg, &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

func GetClients(reg string) ([]Customer, error) {
	var clients []Customer
	db := storage.DB()
	if err := db.GetClients(reg, &clients); err != nil {
		return nil, err
	}
	return clients, nil
}

func InWork() ([]OrderFull, error) {
	var orders []Order
	db := storage.DB()
	if err := db.GetInWorkOrders(&orders); err != nil {
		return nil, err
	}
	var result []OrderFull
	for _, v := range orders {
		var client Customer
		if err := db.GetById(config.Cfg.DB.Coll.Customers, &client, v.CustomerID); err != nil {
			return nil, err
		}
		result = append(result, OrderFull{client, v})
	}
	return result, nil
}

func LoadCache() error {
	db := storage.DB()
	orders, err := InWork()
	if err != nil {
		return err
	}
	for _, v := range orders {
		storage.Cache.Update(config.Cfg.DB.Coll.Orders, v.Order.ID, &v)
	}
	var cargos []Cargo
	if err := db.GetAll(config.Cfg.DB.Coll.Cargos, &cargos); err != nil {
		return err
	}
	for _, v := range cargos {
		storage.Cache.Update(config.Cfg.DB.Coll.Cargos, v.ID, v)
	}
	var chemistry []Chemistry
	if err := db.GetAll(config.Cfg.DB.Coll.Chemistry, &chemistry); err != nil {
		return err
	}
	for _, v := range chemistry {
		storage.Cache.Update(config.Cfg.DB.Coll.Chemistry, v.ID, v)
	}
	var users []User
	if err := db.GetAll(config.Cfg.DB.Coll.Users, &users); err != nil {
		return err
	}
	for _, v := range users {
		storage.Cache.Update(config.Cfg.DB.Coll.Users, v.ID, v)
	}
	return nil
}
