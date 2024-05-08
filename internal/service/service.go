package service

type IDBPage interface {
	NewID(string)
	GetID() string
	GetName() string
}
