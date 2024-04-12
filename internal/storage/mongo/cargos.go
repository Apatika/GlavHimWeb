package mongo

import (
	"context"
	"glavhim-app/internal/service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoDB) GetCargos() ([]service.Cargo, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return nil, err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection("cargos")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	var results []service.Cargo
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *MongoDB) AddCargo(cargo service.Cargo) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection("cargos")
	if _, err := coll.InsertOne(context.TODO(), cargo); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdateCargo(cargo service.Cargo) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection("cargos")
	if _, err := coll.UpdateOne(context.TODO(), bson.M{"_id": cargo.ID}, bson.M{"$set": cargo}); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) DeleteCargo(cargo service.Cargo) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection("cargos")
	if _, err := coll.DeleteOne(context.TODO(), bson.M{"_id": cargo.ID}); err != nil {
		return err
	}
	return nil
}
