package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var contactCollectionName = "contact"
var contactUnique = "name"

type Contact struct {
	Text    string `json:"text"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Maps    string `json:"maps"`
	Email   string `json:"email"`
	Telefon string `json:"telefon"`
	Navn    string `json:"navn"`
	Emne    string `json:"emne"`
}

func mapToContact(dataMap map[string]any) (*Contact, error) {
	var contact Contact

	// Check if required keys exist
	if dataMap["name"] == nil || dataMap["age"] == nil || dataMap["img"] == nil {
		return nil, errors.New("missing required key in dataMap")
	}

	return &contact, nil
}

func (contact Contact) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, contactCollectionName, contactUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (contact Contact) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, contactCollectionName, contactUnique)
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

func (contact Contact) Post(data map[string]any) error {
	mappedData, err := mapToContact(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, contactCollectionName, contactUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (contact Contact) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToContact(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, contactCollectionName, contactUnique)
	defer store.Close()

	filter := bson.M{
		"_id": objectID,
	}

	updateData := bson.M{
		"$set": mappedData,
	}

	err = store.UpdateData(filter, updateData)

	return err
}

func (contact Contact) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, contactCollectionName, contactUnique)
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
