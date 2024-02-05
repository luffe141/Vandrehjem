package models

type AboutUs struct {
	Img   string `json:"img"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (u AboutUs) Validate(data any) error {
	// no validation of data
	return nil
}

func (AboutUs) GetCollectionName() string {
	return "about_us"
}

func (AboutUs) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var AboutUsModel = &Model{&AboutUs{}}
