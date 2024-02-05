package models

import "errors"

type AboutUs struct {
	Img   string `json:"img"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (u AboutUs) Validate(a any) error {
	//TODO implement me
	return errors.New("not implemented")
}

func (AboutUs) GetCollectionName() string {
	return "about_us"
}

func (AboutUs) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var AboutUsModel = &Model{&AboutUs{}}
