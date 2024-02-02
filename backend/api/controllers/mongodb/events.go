package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var eventCollectionName = "event"
var eventUnique = "name"

type Events struct {
	//	_Id  string

	Image    string   `json:"image"`
	Images   []string `json:"images"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Contents []string `json:"contents"`
}

func mapToEvents(dataMap map[string]interface{}) (*Events, error) {
	var event Events

	if val, ok := dataMap["image"].(string); ok {
		event.Image = val
	} else {
		return nil, errors.New("image field not found or is not a string")
	}

	if val, ok := dataMap["images"].([]string); ok {
		event.Images = val
	} else {
		return nil, errors.New("images field not found or is not a slice of string")
	}

	if val, ok := dataMap["title"].(string); ok {
		event.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["content"].(string); ok {
		event.Content = val
	} else {
		return nil, errors.New("content field not found or is not a string")
	}

	if val, ok := dataMap["contents"].([]string); ok {
		event.Contents = val
	} else {
		return nil, errors.New("contents field not found or is not a slice of string")
	}

	return &event, nil
}

func (event Events) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, eventCollectionName, eventUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (event Events) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, eventCollectionName, eventUnique)
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

func (events Events) Post(data map[string]interface{}) error {
	mappedData, err := mapToEvents(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ReviewCollectionName, ReviewUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (events Events) Put(id string, data map[string]interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToEvents(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ReviewCollectionName, ReviewUnique)
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

func (event Events) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, eventCollectionName, eventUnique)
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
