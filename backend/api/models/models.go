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

func (s *Model) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, s.GetCollectionName(), s.GetUnique())
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (s *Model) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, s.GetCollectionName(), s.GetUnique())
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

func (s *Model) Post(data any) error {
	err := s.Validate(data)
	if err != nil {
		return err
	}

	// convert your struct back to map here
	// as you transform your json into a struct in handler
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, s.GetCollectionName(), s.GetUnique())
	defer store.Close()

	err = store.CreateData(data) // data is a struct now, not a map
	return err
}

func (s *Model) Put(id string, data map[string]any) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, s.GetCollectionName(), s.GetUnique())
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

func (s *Model) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, s.GetCollectionName(), s.GetUnique())
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
