package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func SetDB(mdb *mongo.Client) {
	db = mdb
}

type Mongo struct {
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("product").Collection(collectionName)
	return collection
}

func (m Mongo) InsertOne(ctx context.Context, document interface{}, collectionName string) (string, error) {
	var collection *mongo.Collection = GetCollection(db, collectionName)
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}
	return primitive.ObjectID.Hex(result.InsertedID.(primitive.ObjectID)), nil
}

func (m Mongo) FindById(ctx context.Context, id string, collectionName string, result interface{}) error {
	var collection *mongo.Collection = GetCollection(db, collectionName)
	objId, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (m Mongo) Find(ctx context.Context, collectionName string, filter interface{}, skip int, limit int, result interface{}) error {
	var collection *mongo.Collection = GetCollection(db, collectionName)
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))
	fmt.Println(limit)
	fmt.Println(skip)
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		fmt.Println(err)
	}
	if err = cursor.All(context.TODO(), result); err != nil {
		return err
	}
	return nil
}
