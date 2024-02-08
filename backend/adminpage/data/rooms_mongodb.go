package data

import (
	"github.com/lmbek/bekrouter/mvc/controllers/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRooms() (any, error) {
	var CollectionName = "rooms"
	storage := mongodb.NewStorage(MongodbConnection, DatabaseName, CollectionName)
	return storage.ReadData(bson.M{})
}
