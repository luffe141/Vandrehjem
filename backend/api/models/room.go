package models

import (
	"errors"
)

type Room struct {
	//	_Id  string
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Img     string   `json:"img"`
	Price   string   `json:"price"`
	Slider  []string `json:"slider"`
	Content string   `json:"content"`
}

func (r Room) Validate(data any) error {
	// no validation of data
	return nil
}

func (Room) GetCollectionName() string {
	return "room"
}

func (Room) GetUnique() string {
	return "name"
}

var RoomModel = &Model{&Room{}}

func mapToRoom(dataMap map[string]interface{}) (*Room, error) {
	var room Room

	if val, ok := dataMap["name"].(string); ok {
		room.Name = val
	} else {
		return nil, errors.New("name field not found or is not a string")
	}

	if val, ok := dataMap["title"].(string); ok {
		room.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["text"].(string); ok {
		room.Text = val
	} else {
		return nil, errors.New("text field not found or is not a string")
	}

	if val, ok := dataMap["img"].(string); ok {
		room.Img = val
	} else {
		return nil, errors.New("img field not found or is not a string")
	}

	if val, ok := dataMap["price"].(string); ok {
		room.Price = val
	} else {
		return nil, errors.New("price field not found or is not a string")
	}

	if val, ok := dataMap["slider"].([]string); ok {
		room.Slider = val
	} else {
		return nil, errors.New("slider field not found or is not a slice of string")
	}

	if val, ok := dataMap["content"].(string); ok {
		room.Content = val
	} else {
		return nil, errors.New("content field not found or is not a string")
	}

	return &room, nil
}
