package service

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
	"log"
)

type Cargo struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	URI        string `json:"uri" bson:"uri"`
	MainTel    string `json:"mainTel" bson:"mainTel"`
	ManagerTel string `json:"managerTel" bson:"managerTel"`
}

func (c *Cargo) GetAll() (map[string]interface{}, error) {
	var cargos []Cargo
	result := make(map[string]interface{}, 1)
	if err := storage.DB.GetAll(config.Cfg.DB.Coll.Cargos, &cargos); err != nil {
		return nil, err
	}
	//костыль, можно убрать, если переделать фронтенд под массивы
	for _, v := range cargos {
		result[v.ID] = v
	}
	return result, nil
}

func (c *Cargo) Push(raw []byte) error {
	json.Unmarshal(raw, c)
	c.ID = storage.DB.GetNewID()
	if err := storage.DB.CheckName(c.Name, config.Cfg.DB.Coll.Cargos, c.ID); err != nil {
		return fmt.Errorf("write cargo to db failed(%v)", err.Error())
	}
	if err := storage.DB.Add(config.Cfg.DB.Coll.Cargos, c); err != nil {
		return fmt.Errorf("write cargo to db failed(%v)", err.Error())
	}
	log.Printf("create cargo (id: %v)", c.ID)
	return nil
}

func (c *Cargo) Update(raw []byte) error {
	json.Unmarshal(raw, c)
	if err := storage.DB.CheckName(c.Name, config.Cfg.DB.Coll.Cargos, c.ID); err != nil {
		return fmt.Errorf("update cargo failed(%v)", err.Error())
	}
	if err := storage.DB.Update(config.Cfg.DB.Coll.Cargos, c, c.ID); err != nil {
		return fmt.Errorf("update cargo failed(%v)", err.Error())
	}
	log.Printf("update cargo (id: %v)", c.ID)
	return nil
}

func (c *Cargo) Delete(raw []byte) error {
	json.Unmarshal(raw, c)
	if err := storage.DB.Delete(config.Cfg.DB.Coll.Cargos, c.ID); err != nil {
		return fmt.Errorf("delete cargo from db failed (%v)", err.Error())
	}
	log.Printf("delete cargo (name: %v)", c.Name)
	return nil
}
