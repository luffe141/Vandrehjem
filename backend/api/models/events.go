package models

type Events struct {
	//	_Id  string
	Image    string   `json:"image"`
	Images   []string `json:"images"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Contents []string `json:"contents"`
}

func (e Events) Validate(data any) error {
	// no validation of data
	return nil
}

func (Events) GetCollectionName() string {
	return "event"
}

func (Events) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var EventsModel = &Model{&Events{}}
