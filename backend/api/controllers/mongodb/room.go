package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

var roomCollectionName = "room"
var roomUnique = "name"

type Room struct {
	//	_Id  string
	Name  string `json:"name"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Img   string `json:"img"`
	Price string `json:"price"`
	Slider []string `json:"slider"`
	Content string `json:"content"`
}

func mapToRoom(dataMap map[string]any) (*Room, error) {
	var room Room

	// Check if required keys exist
	if dataMap["name"] == nil || dataMap["age"] == nil || dataMap["img"] == nil {
		return nil, errors.New("missing required key in dataMap")
	}

	// Iterate over provided map
	for key, value := range dataMap {
		switch strings.ToLower(key) {
		case "name":
			name, ok := value.(string)
			if !ok || name == "" {
				return nil, errors.New("invalid or empty Name field")
			}
			room.Name = name
		case "age":
			str := value.(string)
			age, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}

			if age == 0 {
				age = -1
			}

			room.Age = age
		case "img":
			img, ok := value.(string)
			if !ok || img == "" {
				return nil, errors.New("invalid or empty Img field")
			}
			room.Img = img
		}
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

func (room Room) Post(data map[string]any) error {
	mappedData, err := mapToRoom(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (room Room) Put(id string, data map[string]any) error {
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

	err = store.UpdateData(filter, updateData)

	return err
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
