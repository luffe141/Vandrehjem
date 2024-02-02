package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var roomCollectionName = "room"
var roomUnique = "name"

type Room struct {
	//	_Id  string
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Img     string   `json:"img"`
	Price   string   `json:"price"`
	Slider  []string `json:"slider"`
	Content string   `json:"content"`
}

func mapToRoom(dataMap map[string]interface{}) (*Room, error) {
	var room Room

	if val, ok := dataMap["name"].(string); ok {
		room.Name = val
	} else {
		return nil, errors.New("name field not found or is not a string")
	}

	if val, ok := dataMap["title"].(string); ok {
		room.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["text"].(string); ok {
		room.Text = val
	} else {
		return nil, errors.New("text field not found or is not a string")
	}

	if val, ok := dataMap["img"].(string); ok {
		room.Img = val
	} else {
		return nil, errors.New("img field not found or is not a string")
	}

	if val, ok := dataMap["price"].(string); ok {
		room.Price = val
	} else {
		return nil, errors.New("price field not found or is not a string")
	}

	if val, ok := dataMap["slider"].([]string); ok {
		room.Slider = val
	} else {
		return nil, errors.New("slider field not found or is not a slice of string")
	}

	if val, ok := dataMap["content"].(string); ok {
		room.Content = val
	} else {
		return nil, errors.New("content field not found or is not a string")
	}

	return &room, nil
}

func (room Room) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (room Room) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
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

func (room Room) Post(data map[string]interface{}) error {
	mappedData, err := mapToRoom(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
	defer store.Close()

	return store.CreateData(*mappedData)
}

func (room Room) Put(id string, data map[string]interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToRoom(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
	defer store.Close()

	filter := bson.M{
		"_id": objectID,
	}

	updateData := bson.M{
		"$set": mappedData,
	}

	return store.UpdateData(filter, updateData)
}

func (room Room) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
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
