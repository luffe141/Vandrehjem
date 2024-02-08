package models

import (
	"backend/api/mvc/models/utility"
	"github.com/lmbek/bekrouter/mvc/models"
)

type Activity struct {
	Image    string   `bson:"image" json:"image" xml:"image"`
	Images   []string `bson:"images" json:"images" xml:"images"`
	Title    string   `bson:"title" json:"title" xml:"title"`
	Titles   string   `bson:"titles" json:"titles" xml:"titles"`
	Content  string   `bson:"content" json:"content" xml:"content"`
	Distance string   `bson:"distance" json:"distance" xml:"distance"`
}

func (activity *Activity) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	//requiredKeys := []string{"title", "content", "distance"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*Activity) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*Activity) GetCollectionName() string {

	return "activities"
}

func (*Activity) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var ActivityModel = &models.Model{IModel: &Activity{}, MongodbConnection: MongodbConnection, DatabaseName: DatabaseName}
