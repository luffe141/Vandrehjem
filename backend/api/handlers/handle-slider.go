package api

import (
	"backend/database/mongodb"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Slider struct {
	Name string
	Age  int
	img  string
}

func handleGetSlider(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Slider", "name")
	defer store.Close()

	data, err := store.ReadData(bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("hej med dig")
	fmt.Println(data)
	_, _ = fmt.Fprint(response, data)
}

func handlePutSlider(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Slider", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Slider{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func hadnlePostSlider(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Slider", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Slider{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func handelDeleteSlider(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Slider", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Slider{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}
