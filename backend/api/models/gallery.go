package models

type Gallery struct {
	//	_Id  string
	Title  string   `json:"title"`
	Images []string `json:"images"`
}

func (g Gallery) Validate(a any) error {
	//TODO implement me
	panic("implement me")
}

func (Gallery) GetCollectionName() string {
	return "gallery"
}

func (Gallery) GetUnique() string {
	return "" // From your initial code, activity does not have a unique field.
}

var GalleryModel = &Model{&Gallery{}}
