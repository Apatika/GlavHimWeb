package service

type Pvd struct {
	Weight float32
	Count  int
}

type Order struct {
	CreationDay   int
	CreationMonth int
	CreationYear  int
	LastDate      string
	Cargo         string
	ClientID      int
	Invoice       []int
	Comment       string
	Probes        []Chemistry
	Pvd           []Pvd
	Payment       bool
	ShipmentDay   int
	ShipmentMonth int
	ShipmentYear  int
	Status        int
}
