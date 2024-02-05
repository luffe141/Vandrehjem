package models

import "errors"

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

func (c Contact) Validate(a any) error {
	//TODO implement me
	return errors.New("not implemented")
}

func (Contact) GetCollectionName() string {
	return "contact"
}

func (Contact) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var ContactModel = &Model{&Contact{}}
