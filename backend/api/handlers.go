package api

import (
	"backend/database/mongodb"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Users struct {
	Name string
	Age  int
}

func handleGetUsers(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Users", "name")
	defer store.Close()

	data, err := store.ReadData(bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("hej med dig")
	fmt.Println(data)
	_, _ = fmt.Fprint(response, data)
}

func handlePostUsers(response http.ResponseWriter, request *http.Request) {
	store := mongodb.NewStorage("mongodb://localhost:27017", "Vandrerhjem", "Users", "name")
	defer store.Close()

	// Create
	err := store.CreateData(Users{Name: "Lars", Age: 29})
	if err != nil {
		_, _ = fmt.Fprint(response, err)
	} else {
		_, _ = fmt.Fprint(response, "you have succesfully added a new user")
	}

}

func handleGetUsersa(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "her er alle brugerne 2")
}

func handleGetUsersb(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "her er alle brugerne 3")
}

func handleGetUsersId(response http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	fmt.Fprint(response, "her er bruger "+id)
}
