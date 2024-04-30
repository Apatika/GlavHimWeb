package storage

import (
	"glavhim-app/internal/storage/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDataBase interface {
	HealthCheck() error
	CheckNameOne(string, string, primitive.ObjectID) error
	GetAll(string) ([]bson.M, error)
	AddOne(string, interface{}) error
	UpdateOne(string, interface{}, primitive.ObjectID) error
	DeleteOne(string, primitive.ObjectID) error
	GetInWorkOrders() ([]mongodb.OrderFull, error)
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

func CheckNameOne(name string, collName string, id primitive.ObjectID) error {
	if err := db.CheckNameOne(name, collName, id); err != nil {
		return err
	}
	return nil
}

func GetAll(path string) ([]bson.M, error) {
	var data []bson.M
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

func UpdateOne(collName string, obj interface{}, id primitive.ObjectID) error {
	if err := db.UpdateOne(collName, obj, id); err != nil {
		return err
	}
	return nil
}

func DeleteOne(collName string, id primitive.ObjectID) error {
	if err := db.DeleteOne(collName, id); err != nil {
		return err
	}
	return nil
}

func GetInWorkOrders() ([]mongodb.OrderFull, error) {
	data, err := db.GetInWorkOrders()
	if err != nil {
		return nil, err
	}
	return data, nil
}
