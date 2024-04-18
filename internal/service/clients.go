package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Type           int                `json:"type" bson:"type"`
	Manager        User               `json:"manager" bson:"manager"`
	Inn            int64              `json:"inn" bson:"inn"`
	PassportSerial int                `json:"passportSerial" bson:"passport_serial"`
	PassportNum    int                `json:"passportNum" bson:"passport_num"`
	Name           string             `json:"name" bson:"name"`
	Surname        string             `json:"surname" bson:"surname"`
	SecondName     string             `json:"secondName" bson:"second_name"`
	Adress         Adress             `json:"adress" bson:"adress"`
	Contact        []Contact          `json:"contact" bson:"contact"`
	Email          string             `json:"email" bson:"email"`
}
