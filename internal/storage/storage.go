package storage

type IDataBase interface {
	Ping() error
}

func New() IDataBase {
	return NewSQLite()
}
