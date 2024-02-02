package api

import (
	"fmt"
	"net/http"
)

func RouteDontExist(response http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintln(response, "Route dont exist, remember to end your link with a slash /, also look up the routing manual")
}

func RouteMalformed(response http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintln(response, "Route dont exist, did you forget to put in an identifier?")
}

func RouteDontExistMissingSlash(response http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintln(response, "Route dont exist, remember to end your link with a slash /")
}
