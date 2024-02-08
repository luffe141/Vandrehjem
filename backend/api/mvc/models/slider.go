package models

import (
	"backend/api/mvc/models/utility"
	"github.com/lmbek/bekrouter/mvc/models"
)

type Slider struct {
	//	_Id  string
	Name   string   `json:"name"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Slider []string `json:"slider"`
	Img    string   `json:"img"`
}

func (*Slider) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*Slider) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*Slider) GetCollectionName() string {

	return "sliders"
}

func (*Slider) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var SliderModel = &models.Model{IModel: &Slider{}, MongodbConnection: MongodbConnection, DatabaseName: DatabaseName}
