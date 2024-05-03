package storage

import (
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/mongodb"
)

type IDataBase interface {
	HealthCheck() error
	CheckNameOne(string, string, string) error
	GetAll(string) (interface{}, error)
	AddOne(string, interface{}) error
	UpdateOne(string, interface{}, string) error
	DeleteOne(string, string) error
	GetInWorkOrders() ([]service.OrderFull, error)
	GetNewID() string
}

var db IDataBase

func DBInit() {
	db = mongodb.NewMongo()
}

func HealthCheck() error {
	if err := db.HealthCheck(); err != nil {
		return err
	}
	return nil
}

func GetNewID() string {
	return db.GetNewID()
}

func CheckNameOne(name string, collName string, id string) error {
	if err := db.CheckNameOne(name, collName, id); err != nil {
		return err
	}
	return nil
}

func GetAll(path string) (interface{}, error) {
	data, err := db.GetAll(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AddOne(collName string, obj interface{}) error {
	if err := db.AddOne(collName, obj); err != nil {
		return err
	}
	return nil
}

func UpdateOne(collName string, obj interface{}, id string) error {
	if err := db.UpdateOne(collName, obj, id); err != nil {
		return err
	}
	return nil
}

func DeleteOne(collName string, id string) error {
	if err := db.DeleteOne(collName, id); err != nil {
		return err
	}
	return nil
}

func getInWorkOrders() ([]service.OrderFull, error) {
	data, err := db.GetInWorkOrders()
	if err != nil {
		return nil, err
	}
	return data, nil
}
