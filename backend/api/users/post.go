package users

import (
	"fmt"
	"net/http"
)

func Post(response http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(response, "created new task\n")
}
