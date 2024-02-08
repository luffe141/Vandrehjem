package models

import (
	"backend/api/mvc/models/utility"
	"github.com/lmbek/bekrouter/mvc/models"
)

// TODO: test if json field thing is even needed
type Contact struct {
	Text    string `json:"text"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Maps    string `json:"maps"`
	Email   string `json:"email"`
	Telefon string `json:"telefon"`
	Navn    string `json:"navn"`
	Emne    string `json:"emne"`
}

func (contact *Contact) CheckRequired(data map[string]interface{}) error {
	requiredKeys := []string{"title"}
	return utility.EnsureKeys(requiredKeys, data)
}

func (*Contact) Validate(data map[string]any) error {
	// title validation
	//if strValue, ok := data["image"].(string); ok {
	//	if len(strValue) < 5 || len(strValue) > 30 {
	//		return errors.New("The image should be between 5 and 30 characters long")
	//	}
	//}

	// Add additional validation logic for the other keys as needed
	return nil
}

func (*Contact) GetCollectionName() string {

	return "contacts"
}

func (*Contact) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var ContactModel = &models.Model{IModel: &Contact{}, MongodbConnection: MongodbConnection, DatabaseName: DatabaseName}
