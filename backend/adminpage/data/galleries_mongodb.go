package data

import (
	"backend/api/mvc/models"
	"encoding/json"
	"fmt"
	"github.com/lmbek/bekrouter/mvc/controllers/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func GetGalleries() (models.Gallery, error) {
	var MongodbConnection = "mongodb://localhost:27017"
	var DatabaseName = "vandrerhjem"

	// Assuming NewStorage properly sets up a MongoDB connection and returns a type with a ReadData method
	storage := mongodb.NewStorage(MongodbConnection, DatabaseName, "galleries")
	fmt.Println("Attempting to read data...")
	data, err := storage.ReadData(bson.M{}) // bson.M{} to fetch all documents; adjust the filter as necessary
	if err != nil {
		fmt.Println("Error reading data:", err)
		return models.Gallery{}, err
	}

	var gallery models.Gallery
	// Convert []bson.M to JSON []byte
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
	}

	json.Unmarshal(jsonBytes, &gallery)

	fmt.Println("Successfully fetched gallery:", gallery)
	return gallery, nil
}
