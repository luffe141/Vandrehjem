package mongodb

import (
	"backend/database/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

var SliderCollectionName = "slider"
var SliderUnique = "name"

type Slider struct {
	//	_Id  string
	Name  string `json:"name"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Age   int    `json:"age"`
	Img   string `json:"img"`
}

func mapToSlider(dataMap map[string]any) (*Slider, error) {
	var Slider Slider

	// Check if required keys exist
	if dataMap["name"] == nil || dataMap["age"] == nil || dataMap["img"] == nil {
		return nil, errors.New("missing required key in dataMap")
	}

	// Iterate over provided map
	for key, value := range dataMap {
		switch strings.ToLower(key) {
		case "name":
			name, ok := value.(string)
			if !ok || name == "" {
				return nil, errors.New("invalid or empty Name field")
			}
			Slider.Name = name
		case "age":
			str := value.(string)
			age, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}

			if age == 0 {
				age = -1
			}

			Slider.Age = age
		case "img":
			img, ok := value.(string)
			if !ok || img == "" {
				return nil, errors.New("invalid or empty Img field")
			}
			Slider.Img = img
		}
	}

	return &Slider, nil
}

func (Slider Slider) GetAll() (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
	defer store.Close()
	return store.ReadData(bson.M{})
}

func (Slider Slider) GetById(id string) (any, error) {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
	defer store.Close()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
	}

	return store.ReadData(filter)
}

func (Slider Slider) Post(data map[string]any) error {
	mappedData, err := mapToSlider(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
	defer store.Close()

	err = store.CreateData(*mappedData)

	return err
}

func (Slider Slider) Put(id string, data map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mappedData, err := mapToSlider(data)
	if err != nil {
		return err
	}

	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
	defer store.Close()

	filter := bson.M{
		"_id": objectID,
	}

	updateData := bson.M{
		"$set": mappedData,
	}

	err = store.UpdateData(filter, updateData)

	return err
}

func (slider Slider) Delete(id string) error {
	store := mongodb.NewStorage(MongodbConnection, DatabaseName, SliderCollectionName, SliderUnique)
	defer store.Close()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objectID,
	}

	return store.DeleteData(filter)
}
