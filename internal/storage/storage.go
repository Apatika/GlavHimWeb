package storage

import (
	"glavhim-app/internal/storage/mongodb"
)

type IDataBase interface {
	CheckName(name string, collName string, id string) error
	GetById(collName string, obj interface{}, id string) error
	GetAll(collName string, obj interface{}) error
	Add(collName string, obj interface{}) error
	Update(collName string, obj interface{}, id string) error
	Delete(collName string, id string) error
	GetInWorkOrders(obj interface{}) error
	GetNewID() string
	CheckClient(params ...string) (string, error)
	GetCities(reg string, cities interface{}) error
	GetCustomers(reg string, clients interface{}) error
}

func DB() IDataBase {
	return mongodb.NewMongo()
}
