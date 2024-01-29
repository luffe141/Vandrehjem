package users

import (
	"fmt"
	"net/http"
)

// Put - Updates the task by id provided
func Put(response http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	_, _ = fmt.Fprintf(response, "updating task with id=%v\n", id)
}
