package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AboutUsCollectionName = "about us"
var AboutUsUnique = "name"

type AboutUs struct {
	//	_Id  string
	Img   string `json:"img"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func mapToAboutUs(dataMap map[string]interface{}) (*AboutUs, error) {
	var aboutUs AboutUs

	// Make sure all fields are present for type assertion.
	if val, ok := dataMap["img"].(string); ok {
		aboutUs.Img = val
	} else {
		return nil, errors.New("img field not found or is not a string")
	}

	if val, ok := dataMap["title"].(string); ok {
		aboutUs.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["text"].(string); ok {
		aboutUs.Text = val
	} else {
		return nil, errors.New("text field not found or is not a string")
	}

	return &aboutUs, nil
}

func (aboutUs AboutUs) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, AboutUsCollectionName, AboutUsUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (aboutUs AboutUs) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, AboutUsCollectionName, AboutUsUnique)
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

func (aboutUs AboutUs) Post(data map[string]any) error {
	mappedData, err := mapToAboutUs(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, AboutUsCollectionName, AboutUsUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (aboutUs AboutUs) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToAboutUs(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, AboutUsCollectionName, AboutUsUnique)
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

func (aboutUs AboutUs) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, AboutUsCollectionName, AboutUsUnique)
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
