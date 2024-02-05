package models

import (
	"errors"
)

var RestaurantCollectionName = "restaurant"
var RestaurantUnique = "name"

type Restaurant struct {
	//	_Id  string
	Image   string   `json:"image"`
	Images  []string `json:"images"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Menu    string   `json:"menu"`
}

func mapToRestaurant(dataMap map[string]interface{}) (*Restaurant, error) {
	var restaurant Restaurant

	if val, ok := dataMap["image"].(string); ok {
		restaurant.Image = val
	} else {
		return nil, errors.New("image field not found or is not a string")
	}

	if val, ok := dataMap["images"].([]string); ok {
		restaurant.Images = val
	} else {
		return nil, errors.New("images field not found or is not a slice of string")
	}

	if val, ok := dataMap["title"].(string); ok {
		restaurant.Title = val
	} else {
		return nil, errors.New("title field not found or is not a string")
	}

	if val, ok := dataMap["content"].(string); ok {
		restaurant.Content = val
	} else {
		return nil, errors.New("content field not found or is not a string")
	}

	if val, ok := dataMap["menu"].(string); ok {
		restaurant.Menu = val
	} else {
		return nil, errors.New("menu field not found or is not a string")
	}

	return &restaurant, nil
}
