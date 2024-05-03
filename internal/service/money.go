package service

type Money struct {
	ID      string `json:"id" bson:"_id"`
	Day     int    `json:"day" bson:"day"`
	Month   int    `json:"month" bson:"month"`
	Year    int    `json:"year" bson:"year"`
	Price   int    `json:"price" bson:"price"`
	Comment string `json:"comment" bson:"comment"`
}
