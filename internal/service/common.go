package service

import (
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

func SearchOrders(id string, payment bool, month string, limit int64) ([]Order, error) {
	var orders []Order
	if err := storage.DB.SearchOrders(id, payment, month, limit, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
