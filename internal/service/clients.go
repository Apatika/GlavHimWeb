package service

type Contact struct {
	Name string
	Tel  string
}

type Client struct {
	Manager        string
	Inn            int64
	PassportSerial int
	PassportNum    int
	Name           string
	City           string
	Adress         string
	Contact        []Contact
	Email          string
}
