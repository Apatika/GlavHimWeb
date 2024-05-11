package service

import (
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/storage"
	"log"
)

type Client struct {
	ID             string    `json:"id" bson:"_id"`
	Type           string    `json:"type" bson:"type"`
	Manager        string    `json:"manager" bson:"manager"`
	Inn            string    `json:"inn" bson:"inn"`
	PassportSerial string    `json:"passportSerial" bson:"passport_serial"`
	PassportNum    string    `json:"passportNum" bson:"passport_num"`
	Name           string    `json:"name" bson:"name"`
	Surname        string    `json:"surname" bson:"surname"`
	SecondName     string    `json:"secondName" bson:"second_name"`
	Adress         Adress    `json:"adress" bson:"adress"`
	Contact        []Contact `json:"contact" bson:"contact"`
	Email          string    `json:"email" bson:"email"`
}

func (c *Client) Push() error {
	db := storage.DB()
	if err := db.Add(config.Cfg.DB.Coll.Clients, c); err != nil {
		return fmt.Errorf("write client to db failed(%v)", err.Error())
	}
	log.Printf("create client (id: %v)", c.ID)
	return nil
}

func (c *Client) Update() error {
	db := storage.DB()
	if err := db.Update(config.Cfg.DB.Coll.Clients, c, c.ID); err != nil {
		return fmt.Errorf("update client failed(%v)", err.Error())
	}
	log.Printf("update client (id: %v)", c.ID)
	return nil
}

func (c *Client) Delete() error {
	db := storage.DB()
	if err := db.Delete(config.Cfg.DB.Coll.Clients, c.ID); err != nil {
		return fmt.Errorf("delete client from db failed (%v)", err.Error())
	}
	storage.Cache.Delete(config.Cfg.DB.Coll.Clients, c.ID)
	log.Printf("delete client (name: %v)", c.Name)
	return nil
}

func (c *Client) Check() (string, error) {
	db := storage.DB()
	var id string
	var err error
	if c.Inn != "" {
		id, err = db.CheckClient(c.Inn)
	} else if c.PassportNum != "" {
		id, err = db.CheckClient(c.PassportNum, c.PassportSerial)
	} else if c.Surname != "" {
		id, err = db.CheckClient(c.Surname, c.Name, c.SecondName)
	} else {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return id, nil
}
