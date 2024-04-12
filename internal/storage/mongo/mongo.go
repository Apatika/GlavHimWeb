package mongo

import (
	"context"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Uri string
}

func NewMongo() *MongoDB {
	return &MongoDB{
		Uri: config.Cfg.DB.URI,
	}
}

func (m *MongoDB) HealthCheck() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) PushOrder(order service.Order) (err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection("orders")
	result, err := coll.InsertOne(context.TODO(), order)
	if err != nil {
		return err
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
