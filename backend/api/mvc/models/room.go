package models

import (
	"backend/api/mvc/models/utility"
	"github.com/lmbek/bekrouter/mvc/models"
)

type Room struct {
	//	_Id  string
	Name    string   `json:"name" bson:"name"`
	Title   string   `json:"title" bson:"title"`
	Text    string   `json:"text" bson:"text"`
	Img     string   `json:"img" bson:"img"`
	Price   string   `json:"price" bson:"price"`
	Slider  []string `json:"slider" bson:"slider"`
	Content string   `json:"content" bson:"content"`
}

func (*Room) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*Room) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*Room) GetCollectionName() string {

	return "rooms"
}

func (*Room) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var RoomModel = &models.Model{IModel: &Room{}, MongodbConnection: MongodbConnection, DatabaseName: DatabaseName}
