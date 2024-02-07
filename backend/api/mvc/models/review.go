package models

import (
	"backend/api/mvc/models/utility"
	"github.com/lmbek/bekrouter/mvc/models"
)

type Review struct {
	Name     string   `json:"name"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	Img      string   `json:"img"`
	Rating   string   `json:"rating"`
	Content  string   `json:"content"`
	Contents []string `json:"contents"`
}

func (*Review) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*Review) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*Review) GetCollectionName() string {

	return "review"
}

func (*Review) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var ReviewModel = &models.Model{IModel: &Review{}, MongodbConnection: MongodbConnection, DatabaseName: DatabaseName}
