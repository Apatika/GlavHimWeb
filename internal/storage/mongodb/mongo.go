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
	Uri    string
	Client *mongo.Client
}

func NewMongo() *MongoDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Cfg.DB.URI))
	if err != nil {
		log.Panic(err)
	}
	return &MongoDB{
		Uri:    config.Cfg.DB.URI,
		Client: client,
	}
}

func (m *MongoDB) GetNewID() string {
	return primitive.NewObjectID().Hex()
}

func (m *MongoDB) CheckName(name string, collName string, id string) error {
	nc := struct {
		Name string             `bson:"name"`
		ID   primitive.ObjectID `bson:"_id"`
	}{}
	coll := m.Client.Database(config.Cfg.AppName).Collection(collName)
	coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&nc)
	count, err := coll.CountDocuments(context.TODO(), bson.D{{Key: "name", Value: name}}, nil)
	if (count > 0 && nc.Name != name) || (count == 0 && nc.Name == name) {
		return fmt.Errorf("name already exists")
	}
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetAll(collName string, obj interface{}) error {
	coll := m.Client.Database(config.Cfg.AppName).Collection(collName)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), obj); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetById(collName string, obj interface{}, id string) error {
	coll := m.Client.Database(config.Cfg.AppName).Collection(collName)
	if err := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(obj); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Add(collName string, obj interface{}) error {
	coll := m.Client.Database(config.Cfg.AppName).Collection(collName)
	if _, err := coll.InsertOne(context.TODO(), obj); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Update(collName string, obj interface{}, id string) error {
	coll := m.Client.Database(config.Cfg.AppName).Collection(collName)
	if _, err := coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": obj}); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Delete(collName string, id string) error {
	coll := m.Client.Database(config.Cfg.AppName).Collection(collName)
	if _, err := coll.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetInWorkOrders(obj interface{}) error {
	coll := m.Client.Database(config.Cfg.AppName).Collection(config.Cfg.DB.Coll.Orders)
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "status", Value: bson.D{{Key: "$ne", Value: "Отгружен"}}}})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), obj); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) CheckClient(params ...string) (string, error) {
	if len(params) < 1 || len(params) > 3 {
		return "", fmt.Errorf("invalid number of parameters")
	}
	coll := m.Client.Database(config.Cfg.AppName).Collection(config.Cfg.DB.Coll.Customers)
	type id struct {
		ID string `json:"id" bson:"_id"`
	}
	var find id
	var err error
	switch len(params) {
	case 1:
		err = coll.FindOne(
			context.TODO(),
			bson.D{{Key: "inn", Value: params[0]}},
		).Decode(&find)
	case 2:
		err = coll.FindOne(
			context.TODO(),
			bson.D{{Key: "passport_num", Value: params[0]}, {Key: "passport_serial", Value: params[1]}},
		).Decode(&find)
	case 3:
		err = coll.FindOne(
			context.TODO(),
			bson.D{{Key: "surname", Value: params[0]}, {Key: "name", Value: params[1]}, {Key: "second_name", Value: params[2]}},
		).Decode(&find)
	}
	if err != nil {
		return "", err
	}
	return find.ID, nil
}

func (m *MongoDB) GetCities(reg string, cities interface{}) error {
	coll := m.Client.Database(config.Cfg.AppName).Collection(config.Cfg.DB.Coll.City)
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "city", Value: bson.D{{Key: "$regex",
		Value: primitive.Regex{Pattern: fmt.Sprintf("^.*?(%v).*$", reg), Options: "i"}}}}},
		options.Find().SetLimit(8))
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), cities); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetClients(reg string, clients interface{}) error {
	coll := m.Client.Database(config.Cfg.AppName).Collection(config.Cfg.DB.Coll.Customers)
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "$or", Value: bson.A{bson.D{{Key: "surname", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: fmt.Sprintf("^.*?(%v).*$", reg), Options: "i"}}}}},
		bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: fmt.Sprintf("^.*?(%v).*$", reg), Options: "i"}}}}}}}},
		options.Find().SetLimit(8))
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), clients); err != nil {
		return err
	}
	return nil
}
