package service

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
	"log"
)

type User struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Tel   string `json:"tel" bson:"tel"`
	Email string `json:"email" bson:"email"`
}

func (u *User) GetAll() map[string]interface{} {
	return storage.Cache.Get(config.Cfg.DB.Coll.Users)
}

func (u *User) Push(raw []byte) error {
	json.Unmarshal(raw, u)
	u.ID = storage.DB.GetNewID()
	if err := storage.DB.CheckName(u.Name, config.Cfg.DB.Coll.Users, u.ID); err != nil {
		return fmt.Errorf("write user to db failed(%v)", err.Error())
	}
	if err := storage.DB.Add(config.Cfg.DB.Coll.Users, u); err != nil {
		return fmt.Errorf("write user to db failed(%v)", err.Error())
	}
	go storage.Cache.Update(config.Cfg.DB.Coll.Users, u.ID, u)
	log.Printf("create user (id: %v)", u.ID)
	return nil
}

func (u *User) Update(raw []byte) error {
	json.Unmarshal(raw, u)
	if err := storage.DB.CheckName(u.Name, config.Cfg.DB.Coll.Users, u.ID); err != nil {
		return fmt.Errorf("update user failed(%v)", err.Error())
	}
	if err := storage.DB.Update(config.Cfg.DB.Coll.Users, u, u.ID); err != nil {
		return fmt.Errorf("update user failed(%v)", err.Error())
	}
	go storage.Cache.Update(config.Cfg.DB.Coll.Users, u.ID, u)
	log.Printf("update user (id: %v)", u.ID)
	return nil
}

func (u *User) Delete(raw []byte) error {
	json.Unmarshal(raw, u)
	if err := storage.DB.Delete(config.Cfg.DB.Coll.Users, u.ID); err != nil {
		return fmt.Errorf("delete user from db failed (%v)", err.Error())
	}
	go storage.Cache.Delete(config.Cfg.DB.Coll.Users, u.ID)
	log.Printf("delete user (name: %v)", u.Name)
	return nil
}
