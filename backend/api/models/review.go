package models

import (
	"errors"
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

func (Review) GetCollectionName() string {
	return "review"
}

func (Review) GetUnique() string {
	return "name"
}

var ReviewModel = &Model{&Review{}}

func mapToEvent(dataMap map[string]interface{}) (*Events, error) {
	var event Events

	if val, ok := dataMap["image"].(string); ok {
		event.Image = val
	} else {
		return nil, errors.New("image field not found or is not a string")
	}

	if val, ok := dataMap["images"].([]string); ok {
		event.Images = val
	} else {
		return nil, errors.New("images field not found or is not a slice of string")
	}

	if val, ok := dataMap["title"].(string); ok {
		event.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["content"].(string); ok {
		event.Content = val
	} else {
		return nil, errors.New("content field not found or is not a string")
	}

	if val, ok := dataMap["contents"].([]string); ok {
		event.Contents = val
	} else {
		return nil, errors.New("contents field not found or is not a slice of string")
	}

	return &event, nil
}
