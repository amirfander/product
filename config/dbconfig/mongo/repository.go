package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

func SetDB(mdb *mongo.Client) {
	db = mdb
}

type Mongo struct {
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}

func (m Mongo) InsertOne(ctx context.Context, document interface{}, collectionName string) (bool, error) {
	var collection *mongo.Collection = GetCollection(db, collectionName)
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return false, err
	}
	return true, nil
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
