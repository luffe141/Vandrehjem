package models

type Slider struct {
	//	_Id  string
	Name   string   `json:"name"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Slider []string `json:"slider"`
	Img    string   `json:"img"`
}

func (s Slider) Validate(data any) error {
	// no validation of data
	return nil
}

func (Slider) GetCollectionName() string {
	return "slider"
}

func (Slider) GetUnique() string {
	return ""
}

var SliderModel = &Model{&Slider{}}
