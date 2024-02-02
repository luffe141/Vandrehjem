package nope

import (
	"backend/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ActivityCollectionName = "activity"
var ActivityUnique = "name"

type Activity struct {
	_Id  string
	Name string
	Age  int
	Img  string
}

func (activity Activity) GetAll() (any, error) {
	store := mongodb.NewStorage("mongodb://localhost:27017", DatabaseName, ActivityCollectionName, ActivityUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (activity Activity) GetById(id string) (any, error) {
	store := mongodb.NewStorage("mongodb://localhost:27017", DatabaseName, ActivityCollectionName, ActivityUnique)
	defer store.Close()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
	}

	return store.ReadData(filter)
}

func (activity Activity) Post(data any) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ActivityCollectionName, ActivityUnique)
	defer store.Close()

	return store.CreateData(data)
}

func (activity Activity) Put(id string, data any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ActivityCollectionName, ActivityUnique)
	defer store.Close()

	filter := bson.M{
		"_id": objectID,
	}

	updateData := bson.M{
		"$set": data,
	}

	return store.UpdateData(filter, updateData)
}

func (activity Activity) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ActivityCollectionName, ActivityUnique)
	defer store.Close()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objectID,
	}

	return store.DeleteData(filter)
}
