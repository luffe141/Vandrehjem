package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

var AboutUsCollectionName = "about us"
var AboutUsUnique = "name"

type AboutUs struct {
	//	_Id  string
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
	Img  string `json:"img,omitempty"`
}

func mapToAboutUs(dataMap map[string]any) (*AboutUs, error) {
	var aboutUs AboutUs

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
			aboutUs.Name = name
		case "age":
			str := value.(string)
			age, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}

			if age == 0 {
				age = -1
			}

			aboutUs.Age = age
		case "img":
			img, ok := value.(string)
			if !ok || img == "" {
				return nil, errors.New("invalid or empty Img field")
			}
			aboutUs.Img = img
		}
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
