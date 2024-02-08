package models

import (
	"backend/api/mvc/models/utility"
	"github.com/lmbek/bekrouter/mvc/models"
)

type Restaurant struct {
	//	_Id  string
	Image   string   `json:"image"`
	Images  []string `json:"images"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Menu    string   `json:"menu"`
}

func (*Restaurant) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*Restaurant) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*Restaurant) GetCollectionName() string {

	return "restaurant"
}

func (*Restaurant) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var RestaurantModel = &models.Model{IModel: &Restaurant{}, MongodbConnection: MongodbConnection, DatabaseName: DatabaseName}
