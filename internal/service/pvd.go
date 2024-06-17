package service

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
	"log"
)

type Pvd struct {
	ID       string  `json:"id" bson:"_id"`
	Weight   float64 `json:"weight" bson:"weight"`
	Count    int     `json:"count" bson:"count"`
	Reserved int     `json:"reserved" bson:"reserved"`
}

func GetPvds() ([]Pvd, error) {
	var pvds []Pvd
	if err := storage.DB.GetAll(config.Cfg.DB.Coll.Pvd, &pvds); err != nil {
		return nil, err
	}
	return pvds, nil
}

func (p *Pvd) Push(raw []byte) error {
	json.Unmarshal(raw, p)
	p.ID = storage.DB.GetNewID()
	if err := storage.DB.Add(config.Cfg.DB.Coll.Pvd, p); err != nil {
		return fmt.Errorf("write pvd to db failed(%v)", err.Error())
	}
	log.Print("pvd added")
	return nil
}
