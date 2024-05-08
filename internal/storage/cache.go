package storage

import (
	"context"
	"glavhim-app/internal/config"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage/heapcache"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICache interface {
	GetInWork() []service.OrderFull
	NewOrder(service.OrderFull)
	UpdateOrder(service.OrderFull) error
	UpdateOrderStatus(service.OrderStatusChanger) error
	Managers() []service.User
	Cargos() []service.Cargo
	Chemistry() []service.Chemistry
	AddToManagers(service.User)
	AddToCargos(service.Cargo)
	AddToChemistry(service.Chemistry)
	DeleteManagers(service.User)
	DeleteCargos(service.Cargo)
	DeleteChemistry(service.Chemistry)
}

var Cache ICache

func CacheInit() error {
	var err error
	if Cache, err = newCache(); err != nil {
		return err
	}
	return nil
}

func newCache() (ICache, error) {
	db := DB()
	orders, err := db.GetInWorkOrders()
	if err != nil {
		return nil, err
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Cfg.DB.URI))
	if err != nil {
		return nil, err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database(config.Cfg.AppName).Collection(config.Cfg.DB.Coll.Chems)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	var chems = make(map[string]service.Chemistry, 40)
	for cursor.Next(context.TODO()) {
		var result service.Chemistry
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		chems[result.ID] = result
	}
	coll = client.Database(config.Cfg.AppName).Collection(config.Cfg.DB.Coll.Cargos)
	cursor, err = coll.Find(context.TODO(), bson.D{})
	cargos := make(map[string]service.Cargo, 20)
	for cursor.Next(context.TODO()) {
		var result service.Cargo
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		cargos[result.ID] = result
	}
	coll = client.Database(config.Cfg.AppName).Collection(config.Cfg.DB.Coll.Users)
	cursor, err = coll.Find(context.TODO(), bson.D{})
	users := make(map[string]service.User, 8)
	for cursor.Next(context.TODO()) {
		var result service.User
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		users[result.ID] = result
	}
	return heapcache.New(orders, chems, cargos, users), nil
}
