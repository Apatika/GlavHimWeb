package storage

import (
	"glavhim-app/internal/storage/mongodb"
)

type IDataBase interface {
	CheckName(name string, collName string, id string) error
	GetById(collName string, obj interface{}, id string) error
	GetAll(collName string, obj interface{}) error
	GetOrdersByStatus(obj interface{}, status string) error
	GetOrdersByNotStatus(obj interface{}, status string) error
	Add(collName string, obj interface{}) error
	Update(collName string, obj interface{}, id string) error
	Delete(collName string, id string) error
	GetNewID() string
	CheckClient(params ...string) (string, error)
	GetCities(reg string, cities interface{}) error
	GetCustomers(reg string, clients interface{}) error
	SearchOrders(id string, payment bool, month string, limit int64, obj interface{}) error
}

var DB IDataBase

func NewDB() {
	DB = mongodb.NewMongo()
}
