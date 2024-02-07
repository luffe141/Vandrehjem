package models

import (
	"backend/api/mvc/models/utility"
)

type TrainBar struct {
	//	_Id  string
	Name  string `json:"name"`
	Title string `json:"title"`
	Text  string `json:"text"`

	Img string `json:"img"`
}

func (*TrainBar) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*TrainBar) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*TrainBar) GetCollectionName() string {

	return "train bar"
}

func (*TrainBar) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var TrainBarModel = &Model{IModel: &TrainBar{}}
