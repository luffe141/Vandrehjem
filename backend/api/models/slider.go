package models

import (
	"errors"
)

type Slider struct {
	//	_Id  string
	Name   string   `json:"name"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Slider []string `json:"slider"`
	Img    string   `json:"img"`
}

func (Slider) GetCollectionName() string {
	return "slider"
}

func (Slider) GetUnique() string {
	return "name"
}

var SliderModel = &Model{&Slider{}}

func mapToSlider(dataMap map[string]interface{}) (*Slider, error) {
	var slider Slider

	if val, ok := dataMap["name"].(string); ok {
		slider.Name = val
	} else {
		return nil, errors.New("name field not found or is not a string")
	}

	if val, ok := dataMap["title"].(string); ok {
		slider.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["text"].(string); ok {
		slider.Text = val
	} else {
		return nil, errors.New("text field not found or is not a string")
	}

	if val, ok := dataMap["img"].(string); ok {
		slider.Img = val
	} else {
		return nil, errors.New("img field not found or is not a string")
	}

	if val, ok := dataMap["slider"].([]string); ok {
		slider.Slider = val
	} else {
		return nil, errors.New("slider field not found or is not a slice of string")
	}

	return &slider, nil
}
