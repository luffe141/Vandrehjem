package handlers

import (
	"backend/database/mongodb"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Aktiviteter struct {
	Id   string
	Name string
	Age  int
	img  string
}

func HandleGetIdentifierAktiviteter(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Aktiviteter", "name")
	defer store.Close()

	data, err := store.ReadData(bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("hej med dig")
	fmt.Println(data)
	_, _ = fmt.Fprint(response, data)
}

func HandleGetAktiviteter(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Aktiviteter", "name")
	defer store.Close()

	data, err := store.ReadData(bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("hej med dig")
	fmt.Println(data)
	_, _ = fmt.Fprint(response, data)
}

func HandlePutAktiviteter(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Aktiviteter", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Aktiviteter{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func HandlePostAktiviteter(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Aktiviteter", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Aktiviteter{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func HandleDeleteAktiviteter(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Aktiviteter", "name")
	defer store.Close()

	id := request.PathValue("id")

	filter := bson.M{
		"_id": id,
	}
	err := store.DeleteData(filter)
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully deleted a user")
	}

}
