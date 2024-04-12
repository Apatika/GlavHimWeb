package storage

import (
	"glavhim-app/internal/service"
)

type IDataBase interface {
	HealthCheck() error
	PushOrder(service.Order) error
	GetUsers() ([]service.User, error)
	AddUser(service.User) error
	UpdateUser(service.User) error
	DeleteUser(service.User) error
}

var db IDataBase

func DBInit() {
	db = NewMongo()
}

func HealthCheck() error {
	if err := db.HealthCheck(); err != nil {
		return err
	}
	return nil
}

func PushOrder(order service.Order) error {
	if err := db.PushOrder(order); err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]service.User, error) {
	var user []service.User
	user, err := db.GetUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(user service.User) error {
	if err := db.AddUser(user); err != nil {
		return err
	}
	return nil
}

func UpdateUser(user service.User) error {
	if err := db.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func DeleteUser(user service.User) error {
	if err := db.DeleteUser(user); err != nil {
		return err
	}
	return nil
}
