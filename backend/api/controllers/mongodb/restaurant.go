package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var RestaurantCollectionName = "restaurant"
var RestaurantUnique = "name"

type Restaurant struct {
	//	_Id  string
	Image   string   `json:"image"`
	Images  []string `json:"images"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Menu    string   `json:"menu"`
}

func mapToRestaurant(dataMap map[string]interface{}) (*Restaurant, error) {
	var restaurant Restaurant

	if val, ok := dataMap["image"].(string); ok {
		restaurant.Image = val
	} else {
		return nil, errors.New("image field not found or is not a string")
	}

	if val, ok := dataMap["images"].([]string); ok {
		restaurant.Images = val
	} else {
		return nil, errors.New("images field not found or is not a slice of string")
	}

	if val, ok := dataMap["title"].(string); ok {
		restaurant.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["content"].(string); ok {
		restaurant.Content = val
	} else {
		return nil, errors.New("content field not found or is not a string")
	}

	if val, ok := dataMap["menu"].(string); ok {
		restaurant.Menu = val
	} else {
		return nil, errors.New("menu field not found or is not a string")
	}

	return &restaurant, nil
}

func (restaurant Restaurant) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, RestaurantCollectionName, RestaurantUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (restaurant Restaurant) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, RestaurantCollectionName, RestaurantUnique)
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

func (restaurant Restaurant) Post(data map[string]any) error {
	mappedData, err := mapToRestaurant(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, RestaurantCollectionName, RestaurantUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (restaurant Restaurant) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToRestaurant(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, RestaurantCollectionName, RestaurantUnique)
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

func (restaurant Restaurant) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, RestaurantCollectionName, RestaurantUnique)
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
