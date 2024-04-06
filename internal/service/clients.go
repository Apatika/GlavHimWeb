package service

type Contact struct {
	Name string
	Tel  string
}

type Adress struct {
	City   string
	Adress string
}

type Client struct {
	Type           string
	Manager        User
	Inn            int64
	PassportSerial int
	PassportNum    int
	Name           string
	Surname        string
	SecondName     string
	Adress         []Adress
	Contact        []Contact
	Email          string
}
