package storage

type IDataBase interface {
	HealthCheck() error
	PushOrder() error
}
