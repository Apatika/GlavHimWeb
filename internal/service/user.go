package service

type User struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Tel   string `json:"tel" bson:"tel"`
	Email string `json:"email" bson:"email"`
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
