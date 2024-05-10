package storage

import (
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/mongodb"
)

type IDataBase interface {
	HealthCheck() error
	CheckName(string, string, string) error
	GetAll(string, interface{}) error
	Add(string, interface{}) error
	Update(string, interface{}, string) error
	Delete(string, string) error
	GetInWorkOrders() ([]service.OrderFull, error)
	GetNewID() string
	CheckClient(service.Client) (string, error)
	GetCities(string) ([]service.City, error)
}

func DB() IDataBase {
	return mongodb.NewMongo()
}
