package models

import (
	"errors"
	"github.com/lmbek/bekrouter/example/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type IModel interface {
	GetCollectionName() string
	GetUnique() string
	Validate(map[string]any) error

	CheckRequired(map[string]any) error
}

type Model struct {
	IModel
}

func (s *Model) Filter(data map[string]interface{}) map[string]interface{} {
	t := reflect.TypeOf(s.IModel).Elem()
	typeTagMap := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		typeTagMap[jsonTag] = field.Type.Kind().String()
	}

	filteredData := make(map[string]interface{})
	for k, v := range data {
		if _, exists := typeTagMap[k]; exists {
			filteredData[k] = v
		}
	}
	return filteredData
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

func (s *Model) Post(data map[string]any) error {
	data = s.Filter(data)

	err := s.CheckRequired(data)

	// convert your struct back to map here
	// as you transform your json into a struct in handler
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, s.GetCollectionName(), s.GetUnique())
	defer store.Close()

	err = store.CreateData(data) // data is a struct now, not a map
	return err
}

func toBsonM(data map[string]any) bson.M {
	bsonData := make(bson.M)
	for k, v := range data {
		bsonData[k] = v
	}
	return bsonData
}

// Put TODO: Could be reworked so we look in the database first to confirm if item with id exist, before modifying it
func (s *Model) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	data = s.Filter(data)
	err = s.Validate(data)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objectID,
	}

	updateData := bson.M{
		"$set": toBsonM(data),
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, s.GetCollectionName(), s.GetUnique())
	defer store.Close()

	modified, err := store.UpdateData(filter, updateData)
	if err != nil {
		return err
	}

	if !modified {
		return errors.New("there was nothing modified, maybe the id did not exist")
	}

	return nil
}

//func (s *Model) Put(id string, data map[string]any) error {
//	store := mongodb.NewStorage(MongodbConnection, DatabaseName, s.GetCollectionName(), s.GetUnique())
//	defer store.Close()
//
//	objectID, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return err
//	}
//
//	filter := bson.M{
//		"_id": objectID,
//	}
//
//	// todo: data wants to be bson.M?
//	updateData := bson.M{
//		"$set": data,
//	}
//
//	return store.UpdateData(filter, updateData)
//}

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
