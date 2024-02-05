package models

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var trainBarCollectionName = "train bar"
var trainBarUnique = "name"

type TrainBar struct {
	//	_Id  string
	Name  string `json:"name"`
	Title string `json:"title"`
	Text  string `json:"text"`

	Img string `json:"img"`
}

func (TrainBar) GetCollectionName() string {
	return "train_bar"
}

func (TrainBar) GetUnique() string {
	return "name"
}

var TrainBarModel = &Model{&TrainBar{}}

func mapToTrainBar(dataMap map[string]interface{}) (*TrainBar, error) {
	var trainBar TrainBar

	if val, ok := dataMap["name"].(string); ok {
		trainBar.Name = val
	} else {
		return nil, errors.New("name field not found or is not a string")
	}

	if val, ok := dataMap["title"].(string); ok {
		trainBar.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["text"].(string); ok {
		trainBar.Text = val
	} else {
		return nil, errors.New("text field not found or is not a string")
	}

	if val, ok := dataMap["img"].(string); ok {
		trainBar.Img = val
	} else {
		return nil, errors.New("img field not found or is not a string")
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
