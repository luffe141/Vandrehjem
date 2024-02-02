package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SliderCollectionName = "slider"
var SliderUnique = "name"

type Slider struct {
	//	_Id  string
	Name   string   `json:"name"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Slider []string `json:"slider"`
	Img    string   `json:"img"`
}

func mapToSlider(dataMap map[string]interface{}) (*Slider, error) {
	var slider Slider

	if val, ok := dataMap["name"].(string); ok {
		slider.Name = val
	} else {
		return nil, errors.New("name field not found or is not a string")
	}

	if val, ok := dataMap["title"].(string); ok {
		slider.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["text"].(string); ok {
		slider.Text = val
	} else {
		return nil, errors.New("text field not found or is not a string")
	}

	if val, ok := dataMap["img"].(string); ok {
		slider.Img = val
	} else {
		return nil, errors.New("img field not found or is not a string")
	}

	if val, ok := dataMap["slider"].([]string); ok {
		slider.Slider = val
	} else {
		return nil, errors.New("slider field not found or is not a slice of string")
	}

	return &slider, nil
}

func (Slider Slider) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (Slider Slider) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
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

func (Slider Slider) Post(data map[string]any) error {
	mappedData, err := mapToSlider(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (Slider Slider) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToSlider(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
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

func (slider Slider) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
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
