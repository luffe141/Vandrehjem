package models

type Review struct {
	Name     string   `json:"name"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	Img      string   `json:"img"`
	Rating   string   `json:"rating"`
	Content  string   `json:"content"`
	Contents []string `json:"contents"`
}

func (r Review) Validate(data any) error {
	// no validation of data
	return nil
}

func (Review) GetCollectionName() string {
	return "review"
}

func (Review) GetUnique() string {
	return ""
}

var ReviewModel = &Model{&Review{}}
