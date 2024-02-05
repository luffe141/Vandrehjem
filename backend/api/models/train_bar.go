package models

type TrainBar struct {
	//	_Id  string
	Name  string `json:"name"`
	Title string `json:"title"`
	Text  string `json:"text"`

	Img string `json:"img"`
}

func (TrainBar) Validate(data any) error {
	// no validation of data
	return nil
}

func (TrainBar) GetCollectionName() string {
	return "train_bar"
}

func (TrainBar) GetUnique() string {
	return ""
}

var TrainBarModel = &Model{&TrainBar{}}
