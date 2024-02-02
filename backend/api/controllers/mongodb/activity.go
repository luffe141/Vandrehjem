package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

var ActivityCollectionName = "activity"
var ActivityUnique = "name"

type Activity struct {
	//	_Id  string
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
	Img  string `json:"img,omitempty"`
}

func mapToActivities(dataMap map[string]any) (*Activity, error) {
	var activity Activity

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
			activity.Name = name
		case "age":
			str := value.(string)
			age, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}

			if age == 0 {
				age = -1
			}

			activity.Age = age
		case "img":
			img, ok := value.(string)
			if !ok || img == "" {
				return nil, errors.New("invalid or empty Img field")
			}
			activity.Img = img
		}
	}

	return &activity, nil
}

func (activity Activity) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ActivityCollectionName, ActivityUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (activity Activity) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ActivityCollectionName, ActivityUnique)
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

func (activity Activity) Post(data map[string]any) error {
	mappedData, err := mapToActivities(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ActivityCollectionName, ActivityUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (activity Activity) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToActivities(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ActivityCollectionName, ActivityUnique)
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

func (activity Activity) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ActivityCollectionName, ActivityUnique)
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
