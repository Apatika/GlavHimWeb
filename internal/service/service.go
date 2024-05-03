package service

type IService interface {
	NewID(string)
	GetID() string
	GetName() string
}

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
