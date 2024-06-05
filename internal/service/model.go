package service

const (
	StatusInWork  = "Принят В Работу"
	StatusShipped = "Отгружен"
	StatusPecom   = "Забор ПЭК"
	StatusCalled  = "Заказан Забор"
	StatusStop    = "СТОП"
	StatusCity    = "Развозка"
	StatusEmpty   = "Нет Товара"
	StatusChanged = "Изменен!"
)

type Pvd struct {
	Weight float32 `json:"weight" bson:"weight"`
	Count  int     `json:"count" bson:"count"`
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

type City struct {
	Region string `json:"region" bson:"region"`
	City   string `json:"city" bson:"city"`
}
