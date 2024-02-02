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

func mapToContact(dataMap map[string]interface{}) (*Contact, error) {
	var contact Contact

	if val, ok := dataMap["text"].(string); ok {
		contact.Text = val
	} else {
		return nil, errors.New("text field not found or is not a string")
	}

	if val, ok := dataMap["title"].(string); ok {
		contact.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["content"].(string); ok {
		contact.Content = val
	} else {
		return nil, errors.New("content field not found or is not a string")
	}

	if val, ok := dataMap["maps"].(string); ok {
		contact.Maps = val
	} else {
		return nil, errors.New("maps field not found or is not a string")
	}

	if val, ok := dataMap["email"].(string); ok {
		contact.Email = val
	} else {
		return nil, errors.New("email field not found or is not a string")
	}

	if val, ok := dataMap["telefon"].(string); ok {
		contact.Telefon = val
	} else {
		return nil, errors.New("telefon field not found or is not a string")
	}

	if val, ok := dataMap["navn"].(string); ok {
		contact.Navn = val
	} else {
		return nil, errors.New("navn field not found or is not a string")
	}

	if val, ok := dataMap["emne"].(string); ok {
		contact.Emne = val
	} else {
		return nil, errors.New("emne field not found or is not a string")
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
