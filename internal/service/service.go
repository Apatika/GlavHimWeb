package service

import "glavhim-app/internal/config"

type IDBPage interface {
	NewID(string)
	GetID() string
	GetName() string
	GetType() string
}

func GetStruct(path string) IDBPage {
	switch path {
	case config.Cfg.DB.Coll.Users:
		return NewUser()
	case config.Cfg.DB.Coll.Cargos:
		return NewCargo()
	case config.Cfg.DB.Coll.Chemistry:
		return NewChemistry()
	default:
		return nil
	}
}
