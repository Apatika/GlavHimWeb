package service

import (
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
)

func GetCities(reg string) ([]City, error) {
	var cities []City
	if err := storage.DB.GetCities(reg, &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

func GetCustomers(reg string) ([]Customer, error) {
	var clients []Customer
	if err := storage.DB.GetCustomers(reg, &clients); err != nil {
		return nil, err
	}
	return clients, nil
}

func InWork() ([]OrderFull, error) {
	var orders []Order
	if err := storage.DB.GetOrdersByNotStatus(&orders, "Отгружен"); err != nil {
		return nil, err
	}
	var result []OrderFull
	for _, v := range orders {
		var client Customer
		if err := storage.DB.GetById(config.Cfg.DB.Coll.Customers, &client, v.CustomerID); err != nil {
			return nil, err
		}
		result = append(result, OrderFull{client, v})
	}
	return result, nil
}

func LoadCache() error {
	orders, err := InWork()
	if err != nil {
		return err
	}
	for _, v := range orders {
		storage.Cache.Update(config.Cfg.DB.Coll.Orders, v.Order.ID, &v)
	}
	return nil
}

func SearchOrders(id string, payment bool, month string, limit int64) ([]OrderFull, error) {
	var orders []Order
	if err := storage.DB.SearchOrders(id, payment, month, limit, &orders); err != nil {
		return nil, err
	}
	var result []OrderFull
	for _, v := range orders {
		var client Customer
		if err := storage.DB.GetById(config.Cfg.DB.Coll.Customers, &client, v.CustomerID); err != nil {
			return nil, err
		}
		result = append(result, OrderFull{client, v})
	}
	return result, nil
}
