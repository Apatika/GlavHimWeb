package mongodb

import (
	"context"
	"fmt"
	"glavhim-app/internal/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *MongoDB) CheckNameOne(name string, collName string, id primitive.ObjectID) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	nc := struct {
		Name string             `bson:"name"`
		ID   primitive.ObjectID `bson:"_id"`
	}{}
	coll := client.Database("glavhim").Collection(collName)
	err = coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&nc)
	count, err := coll.CountDocuments(context.TODO(), bson.D{{Key: "name", Value: name}}, nil)
	log.Println(nc.Name)
	if (count > 0 && nc.Name != name) || (count == 0 && nc.Name == name) {
		return fmt.Errorf("name already exists")
	}
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetAll(path string) ([]bson.M, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return nil, err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection(path)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *MongoDB) AddOne(collName string, obj interface{}) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection(collName)
	if _, err := coll.InsertOne(context.TODO(), obj); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdateOne(collName string, obj interface{}, id primitive.ObjectID) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection(collName)
	if _, err := coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": obj}); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) DeleteOne(collName string, id primitive.ObjectID) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection(collName)
	if _, err := coll.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return err
	}
	return nil
}
