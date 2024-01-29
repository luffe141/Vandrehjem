package handlers

import (
	"backend/database/mongodb"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Contact struct {
	Name string
	Age  int
	img  string
}

func HandleGetContact(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Contact", "name")
	defer store.Close()

	data, err := store.ReadData(bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("hej med dig")
	fmt.Println(data)
	_, _ = fmt.Fprint(response, data)
}

func HandlePutContact(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Contact", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Contact{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func HandlePostContact(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Contact", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Contact{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func HandleDelteContact(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Contact", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Contact{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}
