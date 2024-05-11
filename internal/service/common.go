package service

import (
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
)

const (
	StatusInWork  = "Принят В Работу"
	StatusShipped = "Отгружен"
	StatusPecom   = "Забор ПЭК"
	StatusCalled  = "Заказан Забор"
	StatusStop    = "СТОП"
	StatusCity    = "Развозка"
	StatusEmpty   = "Нет Товара"
	StatusChanged = "Изменен!"
)

type Pvd struct {
	Weight float32 `json:"weight" bson:"weight"`
	Count  int     `json:"count" bson:"count"`
}

type Date struct {
	Day   int    `json:"day" bson:"day"`
	Month string `json:"month" bson:"month"`
	Year  int    `json:"year" bson:"year"`
}

type Contact struct {
	Name string `json:"name" bson:"name"`
	Tel  string `json:"tel" bson:"tel"`
}

type Adress struct {
	City     string `json:"city" bson:"city"`
	Adress   string `json:"adress" bson:"adress"`
	Terminal string `json:"terminal" bson:"terminal"`
}

func GetCities(reg string) ([]City, error) {
	var cities []City
	db := storage.DB()
	if err := db.GetCities(reg, &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

func InWork() ([]OrderFull, error) {
	var orders []Order
	db := storage.DB()
	if err := db.GetInWorkOrders(&orders); err != nil {
		return nil, err
	}
	var result []OrderFull
	for _, v := range orders {
		var client Client
		if err := db.GetById(config.Cfg.DB.Coll.Clients, &client, v.ClientID); err != nil {
			return nil, err
		}
		result = append(result, OrderFull{client, v})
	}
	return result, nil
}

func LoadCache() error {
	db := storage.DB()
	orders, err := InWork()
	if err != nil {
		return err
	}
	for _, v := range orders {
		storage.Cache.Update(config.Cfg.DB.Coll.Orders, v.Order.ID, &v)
	}
	var cargos []Cargo
	if err := db.GetAll(config.Cfg.DB.Coll.Cargos, &cargos); err != nil {
		return err
	}
	for _, v := range cargos {
		storage.Cache.Update(config.Cfg.DB.Coll.Cargos, v.ID, v)
	}
	var chemistry []Chemistry
	if err := db.GetAll(config.Cfg.DB.Coll.Chemistry, &chemistry); err != nil {
		return err
	}
	for _, v := range chemistry {
		storage.Cache.Update(config.Cfg.DB.Coll.Chemistry, v.ID, v)
	}
	var users []User
	if err := db.GetAll(config.Cfg.DB.Coll.Users, &users); err != nil {
		return err
	}
	for _, v := range users {
		storage.Cache.Update(config.Cfg.DB.Coll.Users, v.ID, v)
	}
	return nil
}
