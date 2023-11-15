package storage

type IDataBase interface {
	Ping() error
	Prepare() error
}

func New() IDataBase {
	return NewSQLite()
}
