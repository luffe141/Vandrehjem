package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

var trainBarCollectionName = "train bar"
var trainBarUnique = "name"

type TrainBar struct {
	//	_Id  string
	Name  string `json:"name"`
	Title string `json:"title"`
	Text  string `json:"text"`

	Img   string `json:"img"`
}

func mapToTrainBar(dataMap map[string]any) (*TrainBar, error) {
	var trainBar TrainBar

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
			trainBar.Name = name
		case "age":
			str := value.(string)
			age, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}

			if age == 0 {
				age = -1
			}

			trainBar.Age = age
		case "img":
			img, ok := value.(string)
			if !ok || img == "" {
				return nil, errors.New("invalid or empty Img field")
			}
			trainBar.Img = img
		}
	}

	return &trainBar, nil
}

func (trainBar TrainBar) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, trainBarCollectionName, trainBarUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (trainBar TrainBar) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, trainBarCollectionName, trainBarUnique)
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

func (trainBar TrainBar) Post(data map[string]any) error {
	mappedData, err := mapToTrainBar(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, trainBarCollectionName, trainBarUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (trainBar TrainBar) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToTrainBar(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, trainBarCollectionName, trainBarUnique)
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

func (trainBar TrainBar) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, trainBarCollectionName, trainBarUnique)
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
