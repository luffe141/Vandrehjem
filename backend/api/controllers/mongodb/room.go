package mongodb

import (
	"backend/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var roomCollectionName = "room"
var roomUnique = "name"

type Room struct {
	//	_Id  string
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Img     string   `json:"img"`
	Price   string   `json:"price"`
	Slider  []string `json:"slider"`
	Content string   `json:"content"`
}

func (room Room) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (room Room) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
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

func (room Room) Post(data map[string]any) error {

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
	defer store.Close()

	return store.CreateData(data)
}

func (room Room) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
	defer store.Close()

	filter := bson.M{
		"_id": objectID,
	}

	updateData := bson.M{
		"$set": data,
	}

	err = store.UpdateData(filter, updateData)

	return err
}

func (room Room) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, roomCollectionName, roomUnique)
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
