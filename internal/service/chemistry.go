package service

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
	"log"
)

type Chemistry struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	SellValue  int    `json:"sellValue" bson:"sellValue"`
	ProbeValue int    `json:"probeValue" bson:"probeValue"`
	ProbeCount int    `json:"probeCount" bson:"probeCount"`
}

func (c *Chemistry) GetAll() map[string]interface{} {
	return storage.Cache.Get(config.Cfg.DB.Coll.Chemistry)
}

func (c *Chemistry) Push(raw []byte) error {
	json.Unmarshal(raw, c)
	c.ID = storage.DB.GetNewID()
	if err := storage.DB.CheckName(c.Name, config.Cfg.DB.Coll.Chemistry, c.ID); err != nil {
		return fmt.Errorf("write chemistry to db failed(%v)", err.Error())
	}
	if err := storage.DB.Add(config.Cfg.DB.Coll.Chemistry, c); err != nil {
		return fmt.Errorf("write chemistry to db failed(%v)", err.Error())
	}
	go storage.Cache.Update(config.Cfg.DB.Coll.Chemistry, c.ID, c)
	log.Printf("create chemistry (id: %v)", c.ID)
	return nil
}

func (c *Chemistry) Update(raw []byte) error {
	json.Unmarshal(raw, c)
	if err := storage.DB.CheckName(c.Name, config.Cfg.DB.Coll.Chemistry, c.ID); err != nil {
		return fmt.Errorf("update chemistry failed(%v)", err.Error())
	}
	if err := storage.DB.Update(config.Cfg.DB.Coll.Chemistry, c, c.ID); err != nil {
		return fmt.Errorf("update chemistry failed(%v)", err.Error())
	}
	go storage.Cache.Update(config.Cfg.DB.Coll.Chemistry, c.ID, c)
	log.Printf("update chemistry (id: %v)", c.ID)
	return nil
}

func (c *Chemistry) Delete(raw []byte) error {
	json.Unmarshal(raw, c)
	if err := storage.DB.Delete(config.Cfg.DB.Coll.Chemistry, c.ID); err != nil {
		return fmt.Errorf("delete chemistry from db failed (%v)", err.Error())
	}
	go storage.Cache.Delete(config.Cfg.DB.Coll.Chemistry, c.ID)
	log.Printf("delete chemistry (name: %v)", c.Name)
	return nil
}
