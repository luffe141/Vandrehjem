package users

import (
	"fmt"
	"net/http"
)

func Delete(response http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	_, _ = fmt.Fprintf(response, "deleting task with id=%v\n", id)
}
