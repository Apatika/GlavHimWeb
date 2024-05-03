package mongodb

import (
	"context"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/service"
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

func (m *MongoDB) GetNewID() string {
	return primitive.NewObjectID().Hex()
}

func (m *MongoDB) CheckNameOne(name string, collName string, id string) error {
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

func (m *MongoDB) GetAll(path string) (interface{}, error) {
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

func (m *MongoDB) UpdateOne(collName string, obj interface{}, id string) error {
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

func (m *MongoDB) DeleteOne(collName string, id string) error {
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

func (m *MongoDB) GetInWorkOrders() ([]service.OrderFull, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return nil, err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database("glavhim").Collection("orders")
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "status", Value: bson.D{{Key: "$ne", Value: "Отгружен"}}}})
	if err != nil {
		return nil, err
	}
	var orders []bson.M
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}
	coll = client.Database("glavhim").Collection("clients")
	arr := make([]service.OrderFull, 0, 10)
	for _, v := range orders {
		doc, err := bson.Marshal(v)
		if err != nil {
			return nil, err
		}
		var order service.Order
		bson.Unmarshal(doc, &order)
		var client service.Client
		err = coll.FindOne(
			context.TODO(),
			bson.D{{Key: "_id", Value: order.ClientID}},
		).Decode(&client)
		if err != nil {
			return nil, err
		}
		arr = append(arr, service.OrderFull{Client: client, Order: order})
	}
	return arr, nil
}
