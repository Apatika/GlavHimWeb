package storage

import (
	"glavhim-app/internal/storage/mongodb"
)

type IDataBase interface {
	HealthCheck() error
	CheckName(string, string, string) error
	GetById(collName string, obj interface{}, id string) error
	GetAll(string, interface{}) error
	Add(string, interface{}) error
	Update(string, interface{}, string) error
	Delete(string, string) error
	GetInWorkOrders(obj interface{}) error
	GetNewID() string
	CheckClient(params ...string) (string, error)
	GetCities(string, interface{}) error
	GetClients(reg string, clients interface{}) error
}

func DB() IDataBase {
	return mongodb.NewMongo()
}
