package models

import (
	"backend/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MongodbConnection = "mongodb://localhost:27017"
var DatabaseName = "vandrerhjem"

type IModel interface {
	GetCollectionName() string
	GetUnique() string
	Validate(any) error
}

type Model struct {
	IModel
}

func (model *Model) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, model.GetCollectionName(), model.GetUnique())
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (model *Model) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, model.GetCollectionName(), model.GetUnique())
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

func (model *Model) Post(data any) error {
	// return err if there is any input that is not allowed
	err := model.Validate(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, model.GetCollectionName(), model.GetUnique())
	defer store.Close()

	err = store.CreateData(data) // data is a struct now, not a map
	return err
}

func (model *Model) Put(id string, data any) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, model.GetCollectionName(), model.GetUnique())
	defer store.Close()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objectID,
	}

	// todo: data wants to be bson.M?
	updateData := bson.M{
		"$set": data,
	}

	return store.UpdateData(filter, updateData)
}

func (model *Model) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, model.GetCollectionName(), model.GetUnique())
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
