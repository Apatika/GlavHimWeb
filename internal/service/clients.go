package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Type           int
	Manager        User
	Inn            int64
	PassportSerial int `json:"passportSerial" bson:"passport_serial"`
	PassportNum    int `json:"passportNum" bson:"passport_num"`
	Name           string
	Surname        string
	SecondName     string `json:"secondName" bson:"second_name"`
	City           string
	Adress         string
	Contact        string
	Tel            string
	Email          string
}
