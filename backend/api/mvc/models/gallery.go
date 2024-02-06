package models

import (
	"backend/api/mvc/models/utility"
)

type Gallery struct {
	//	_Id  string
	Title  string   `json:"title"`
	Images []string `json:"images"`
}

func (*Gallery) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*Gallery) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*Gallery) GetCollectionName() string {

	return "gallery"
}

func (*Gallery) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var GalleryModel = &Model{IModel: &Gallery{}}
