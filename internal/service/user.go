package service

import "glavhim-app/internal/config"

type User struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Tel   string `json:"tel" bson:"tel"`
	Email string `json:"email" bson:"email"`
	Type  string `json:"type" bson:"type"`
}

func NewUser() IDBPage {
	return &User{
		Type: config.Cfg.DB.Coll.Users,
	}
}

func (u *User) NewID(s string) {
	u.ID = s
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetType() string {
	return u.Type
}
