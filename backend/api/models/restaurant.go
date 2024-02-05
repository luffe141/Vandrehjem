package models

type Restaurant struct {
	//	_Id  string
	Image   string   `json:"image"`
	Images  []string `json:"images"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Menu    string   `json:"menu"`
}

func (Restaurant) Validate() error {
	// no validation of data
	return nil
}

func (Restaurant) GetCollectionName() string {
	return "restaurant"
}

func (Restaurant) GetUnique() string {
	return ""
}
