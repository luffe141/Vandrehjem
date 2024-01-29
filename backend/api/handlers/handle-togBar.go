package handlers

import (
	"backend/database/mongodb"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type togBar struct {
	Name string
	Age  int
	img  string
}

func HandleGettogBar(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "togBar", "name")
	defer store.Close()

	data, err := store.ReadData(bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("hej med dig")
	fmt.Println(data)
	_, _ = fmt.Fprint(response, data)
}

func HandlePuttogBar(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "togBar", "name")
	defer store.Close()

	// Create
	err := store.CreateData(togBar{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func HandlePosttogBar(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "togBar", "name")
	defer store.Close()

	// Create
	err := store.CreateData(togBar{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func HandleDeltetogBar(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "togBar", "name")
	defer store.Close()

	// Create
	err := store.CreateData(togBar{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}
