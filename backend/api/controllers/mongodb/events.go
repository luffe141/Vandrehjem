package mongodb

import (
	"backend/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var eventCollectionName = "event"
var eventUnique = "name"

type event struct {
	//	_Id  string

	Image    string   `json:"image"`
	Images   []string `json:"images"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Contents []string `json:"contents"`
}

func mapToevent(dataMap map[string]any) (*event, error) {
	var event event

	/*
		// Iterate over provided map
		for key, value := range dataMap {
			switch strings.ToLower(key) {
			case "name":
				name, ok := value.(string)
				if !ok || name == "" {
					return nil, errors.New("invalid or empty Name field")
				}
				event.Name = name
			case "age":
				str := value.(string)
				age, err := strconv.Atoi(str)
				if err != nil {
					return nil, err
				}

				if age == 0 {
					age = -1
				}

				event.Age = age
			case "img":
				img, ok := value.(string)
				if !ok || img == "" {
					return nil, errors.New("invalid or empty Img field")
				}
				event.Img = img
			}
		}
	*/

	return &event, nil
}

func (event event) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, eventCollectionName, eventUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (event event) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, eventCollectionName, eventUnique)
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

func (event event) Post(data map[string]any) error {
	mappedData, err := mapToevent(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, eventCollectionName, eventUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (event event) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToevent(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, eventCollectionName, eventUnique)
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

func (event event) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, eventCollectionName, eventUnique)
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
