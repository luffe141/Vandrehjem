package users

import (
	"fmt"
	"net/http"
)

func GetAll(response http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(response, "got all tasks\n")
}

func Get(response http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	_, _ = fmt.Fprintf(response, "getting task with id=%v\n", id)
}
