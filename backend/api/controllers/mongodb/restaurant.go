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

func mapToRestaurant(dataMap map[string]any) (*Restaurant, error) {
	var restaurant Restaurant

	// Check if required keys exist
	if dataMap["name"] == nil || dataMap["age"] == nil || dataMap["img"] == nil {
		return nil, errors.New("missing required key in dataMap")
	}
	/*
		// Iterate over provided map
		for key, value := range dataMap {
			switch strings.ToLower(key) {
			case "name":
				name, ok := value.(string)
				if !ok || name == "" {
					return nil, errors.New("invalid or empty Name field")
				}
				restaurant.Name = name
			case "age":
				str := value.(string)
				age, err := strconv.Atoi(str)
				if err != nil {
					return nil, err
				}

				if age == 0 {
					age = -1
				}

				restaurant.Age = age
			case "img":
				img, ok := value.(string)
				if !ok || img == "" {
					return nil, errors.New("invalid or empty Img field")
				}
				restaurant.Img = img
			}
		}
	*/

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
