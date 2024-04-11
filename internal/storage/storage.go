package storage

import "glavhim-app/internal/service"

type IDataBase interface {
	HealthCheck() error
	PushOrder(service.Order) error
	GetUsers() ([]service.User, error)
	AddUser(service.User) error
	UpdateUser(service.User) error
	DeleteUser(service.User) error
}
