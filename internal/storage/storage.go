package storage

import (
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/mongodb"
)

type IDataBase interface {
	HealthCheck() error
	CheckNameOne(string, string, string) error
	GetAll(string, interface{}) error
	AddOne(string, interface{}) error
	UpdateOne(string, interface{}, string) error
	DeleteOne(string, string) error
	GetInWorkOrders() ([]service.OrderFull, error)
	GetNewID() string
	CheckClient(service.Client) (string, error)
}

func DB() IDataBase {
	return mongodb.NewMongo()
}
