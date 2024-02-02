package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ReviewCollectionName = "review"
var ReviewUnique = "name"

type Review struct {
	//	_Id  string
	Name     string   `json:"name"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	Img      string   `json:"img"`
	Rating   string   `json:"rating"`
	Content  string   `json:"content"`
	Contents []string `json:"contents"`
}

func mapToReview(dataMap map[string]any) (*Review, error) {
	var review Review

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
			review.Name = name
		case "age":
			str := value.(string)
			age, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}

			if age == 0 {
				age = -1
			}

			review.Age = age
		case "img":
			img, ok := value.(string)
			if !ok || img == "" {
				return nil, errors.New("invalid or empty Img field")
			}
			review.Img = img
		}
	}

	return &review, nil
}

func (review Review) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ReviewCollectionName, ReviewUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (review Review) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ReviewCollectionName, ReviewUnique)
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

func (review Review) Post(data map[string]any) error {
	mappedData, err := mapToReview(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ReviewCollectionName, ReviewUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (review Review) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToReview(data)
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

func (review Review) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, ReviewCollectionName, ReviewUnique)
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
