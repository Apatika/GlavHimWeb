package storage

import (
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/mongo"
)

type IDataBase interface {
	HealthCheck() error
	PushOrder(service.Order) error
	GetUsers() ([]service.User, error)
	AddUser(service.User) error
	UpdateUser(service.User) error
	DeleteUser(service.User) error
	GetCargos() ([]service.Cargo, error)
	AddCargo(service.Cargo) error
	UpdateCargo(service.Cargo) error
	DeleteCargo(service.Cargo) error
}

var db IDataBase

func DBInit() {
	db = mongo.NewMongo()
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

func GetCargos() ([]service.Cargo, error) {
	var cargo []service.Cargo
	cargo, err := db.GetCargos()
	if err != nil {
		return nil, err
	}
	return cargo, nil
}

func AddCargo(cargo service.Cargo) error {
	if err := db.AddCargo(cargo); err != nil {
		return err
	}
	return nil
}

func UpdateCargo(cargo service.Cargo) error {
	if err := db.UpdateCargo(cargo); err != nil {
		return err
	}
	return nil
}

func DeleteCargo(cargo service.Cargo) error {
	if err := db.DeleteCargo(cargo); err != nil {
		return err
	}
	return nil
}
