package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var galleryCollectionName = "gallery"
var galleryUnique = "name"

type Gallery struct {
	//	_Id  string
	Title  string   `json:"title"`
	Images []string `json:"images"`
}

func mapToGallery(dataMap map[string]any) (*Gallery, error) {
	var gallery Gallery

	// Check if required keys exist
	if dataMap["name"] == nil || dataMap["age"] == nil || dataMap["img"] == nil {
		return nil, errors.New("missing required key in dataMap")
	}

	return &gallery, nil
}

func (gallery Gallery) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, galleryCollectionName, galleryUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (gallery Gallery) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, galleryCollectionName, galleryUnique)
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

func (gallery Gallery) Post(data map[string]any) error {
	mappedData, err := mapToGallery(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, galleryCollectionName, galleryUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (gallery Gallery) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToGallery(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, galleryCollectionName, galleryUnique)
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

func (gallery Gallery) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, galleryCollectionName, galleryUnique)
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
