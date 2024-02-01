package api

import (
	"encoding/json"
	"net/http"
)

func jsonWrapper(handler func(http.ResponseWriter, *http.Request) (any, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		// Call the handler function which should return your data and error
		data, err := handler(response, request)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set the output content type to JSON
		response.Header().Set("Content-Type", "application/json")
		// Convert and write your data as JSON to the response
		err = json.NewEncoder(response).Encode(data)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
