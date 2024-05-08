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
	coll := client.Database(config.Cfg.AppName).Collection(collName)
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
	coll := client.Database(config.Cfg.AppName).Collection(path)
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
	coll := client.Database(config.Cfg.AppName).Collection(collName)
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
	coll := client.Database(config.Cfg.AppName).Collection(collName)
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
	coll := client.Database(config.Cfg.AppName).Collection(collName)
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
	coll := client.Database(config.Cfg.AppName).Collection("orders")
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "status", Value: bson.D{{Key: "$ne", Value: "Отгружен"}}}})
	if err != nil {
		return nil, err
	}
	var orders []service.Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}
	coll = client.Database(config.Cfg.AppName).Collection("clients")
	arr := make([]service.OrderFull, 0, 10)
	for _, v := range orders {
		var client service.Client
		err = coll.FindOne(
			context.TODO(),
			bson.D{{Key: "_id", Value: v.ClientID}},
		).Decode(&client)
		if err != nil {
			return nil, err
		}
		arr = append(arr, service.OrderFull{Client: client, Order: v})
	}
	return arr, nil
}

func (m *MongoDB) CheckClient(c service.Client) (string, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		return "", err
	}
	defer func() {
		err = client.Disconnect(context.TODO())
	}()
	coll := client.Database(config.Cfg.AppName).Collection("clients")
	var find service.Client
	if c.Inn != "" {
		err = coll.FindOne(
			context.TODO(),
			bson.D{{Key: "inn", Value: c.Inn}},
		).Decode(&find)
	} else if c.PassportNum != "" {
		err = coll.FindOne(
			context.TODO(),
			bson.D{{Key: "passport_num", Value: c.PassportNum}, {Key: "passport_serial", Value: c.PassportSerial}},
		).Decode(&find)
	} else {
		err = coll.FindOne(
			context.TODO(),
			bson.D{{Key: "surname", Value: c.Surname}, {Key: "name", Value: c.Name}, {Key: "second_name", Value: c.SecondName}},
		).Decode(&find)
	}
	if err != nil {
		return "", err
	}
	return find.ID, nil
}
