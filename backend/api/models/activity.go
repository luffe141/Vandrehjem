package models

import "errors"

type Activity struct {
	Image    string   `json:"image"`
	Images   []string `json:"images"`
	Title    string   `json:"title"`
	Titles   string   `json:"titles"`
	Content  string   `json:"content"`
	Distance string   `json:"distance"`
}

func (Activity) Validate(any) error {
	//TODO implement me
	return errors.New("not implemented")
}

func (Activity) GetCollectionName() string {
	return "activity"
}

func (Activity) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var ActivityModel = &Model{&Activity{}}
