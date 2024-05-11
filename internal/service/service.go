package service

import "glavhim-app/internal/config"

type IDBPage interface {
	GetAll() map[string]interface{}
	Push(raw []byte) error
	Update(raw []byte) error
	Delete(raw []byte) error
}

func GetStruct(path string) IDBPage {
	switch path {
	case config.Cfg.DB.Coll.Users:
		return &User{}
	case config.Cfg.DB.Coll.Cargos:
		return &Cargo{}
	case config.Cfg.DB.Coll.Chemistry:
		return &Chemistry{}
	default:
		return nil
	}
}
