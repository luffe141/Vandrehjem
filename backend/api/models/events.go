package models

import "errors"

type Events struct {
	//	_Id  string
	Image    string   `json:"image"`
	Images   []string `json:"images"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Contents []string `json:"contents"`
}

func (e Events) Validate(a any) error {
	//TODO implement me
	return errors.New("not implemented")
}

func (Events) GetCollectionName() string {
	return "event"
}

func (Events) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var EventsModel = &Model{&Events{}}
