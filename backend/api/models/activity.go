package models

import (
	"errors"
	"fmt"
)

type Activity struct {
	Image    string   `bson:"image" json:"image"`
	Images   []string `bson:"images" json:"images"`
	Title    string   `bson:"title" json:"title"`
	Titles   string   `bson:"titles" json:"titles"`
	Content  string   `bson:"content" json:"content"`
	Distance string   `bson:"distance" json:"distance"`
}

func (Activity) Validate(data any) error {
	activity, ok := data.(Activity)
	if !ok {
		// handle error when data is not of Activity type
		return errors.New("Invalid data passed")
	}
	//activity := data.(Activity)
	fmt.Println(activity)
	// title validation
	if len(activity.Title) < 5 || len(activity.Title) > 30 {
		return errors.New("title should be between 5 and 30 characters long")
	}

	return nil
}

func (Activity) GetCollectionName() string {
	return "activity"
}

func (Activity) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var ActivityModel = &Model{&Activity{}}
