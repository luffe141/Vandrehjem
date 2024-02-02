package mongodb

import (
	"backend/database/mongodb"
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

func mapToSlider(dataMap map[string]any) (*Slider, error) {
	var Slider Slider

	return &Slider, nil
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
