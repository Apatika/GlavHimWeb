package storage

import (
	"context"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HealthCheck() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Cfg.DB.URI))
	if err != nil {
		return err
	}
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

func PushOrder(order service.Order) (err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Cfg.DB.URI))
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
