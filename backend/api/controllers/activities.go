package controllers

import (
	"backend/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetActivities() (any, error) {
	store := mongodb.NewStorage("mongodb://localhost:27017", databaseName, activityCollection, "name")
	defer store.Close()

	return store.ReadData(bson.M{})
}

func GetActivity(id string) (any, error) {
	store := mongodb.NewStorage("mongodb://localhost:27017", databaseName, activityCollection, "name")
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

func PostActivity(data any) error {
	store := mongodb.NewStorage("mongodb://localhost:27017", databaseName, activityCollection, "name")
	defer store.Close()

	return store.CreateData(data)
}

func PutActivity(id string, data any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage("mongodb://localhost:27017", databaseName, activityCollection, "name")
	defer store.Close()

	filter := bson.M{
		"_id": objectID,
	}

	updateData := bson.M{
		"$set": data,
	}

	return store.UpdateData(filter, updateData)
}

func DeleteActivity(id string) error {
	store := mongodb.NewStorage("mongodb://localhost:27017", databaseName, activityCollection, "name")
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
