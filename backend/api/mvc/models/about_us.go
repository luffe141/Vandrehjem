package models

import (
	"backend/api/mvc/models/utility"
	"github.com/lmbek/bekrouter/mvc/models"
)

type AboutUs struct {
	Img   string `json:"img" bson:"img" json:"img" xml:"img"`
	Title string `json:"title" bson:"title" json:"title" xml:"title"`
	Text  string `json:"text" bson:"text" json:"text" xml:"text"`
}

func (AboutUs *AboutUs) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*AboutUs) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*AboutUs) GetCollectionName() string {

	return "about us"
}

func (*AboutUs) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var AboutUsModel = &models.Model{IModel: &AboutUs{}, MongodbConnection: MongodbConnection, DatabaseName: DatabaseName}
