package service

type Pvd struct {
	Weight float32
	Count  int
}

type CreationDate struct {
	Day   int
	Month int
	Year  int
}

type ShipmentDate struct {
	Day   int
	Month int
	Year  int
}

type Order struct {
	CreationDate CreationDate
	LastDate     string
	Cargo        string
	ClientID     int
	Invoice      []int
	Comment      string
	Probes       []Chemistry
	Pvd          []Pvd
	Payment      bool
	ShipmentDate ShipmentDate
	Status       int
}
