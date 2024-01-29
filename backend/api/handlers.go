package api

import (
	"fmt"
	"net/http"
)

func handleGetUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "her er alle brugerne 1")
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
