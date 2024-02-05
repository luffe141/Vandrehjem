package controllers

import (
	"net/http"
)

// IHandler is an interface for handling CRUD operations on a resource.
// It defines methods for retrieving, creating, updating, and deleting data.
type IHandler interface {
	GetAll() (any, error)
	GetById(id string) (any, error)
	Post(data any) error
	Put(id string, data any) error
	Delete(id string) error
}

// HandleGet is a higher-order function that takes an IHandler as input and returns a function that handles a GET request.
func HandleGet(handler IHandler) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(w http.ResponseWriter, req *http.Request) (any, error) {
		return handler.GetAll()
	}
}

// HandleGetById is a higher-order function that takes an IHandler as input and returns a function that handles a GET request by ID.
func HandleGetById(handler IHandler) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(response http.ResponseWriter, request *http.Request) (any, error) {
		id := request.PathValue("id")
		return handler.GetById(id)
	}
}

// HandlePost is a higher-order function that takes an IHandler as input and returns a function that handles a POST request.
func HandlePost(handler IHandler) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(response http.ResponseWriter, request *http.Request) (any, error) {
		httpRequest := &HttpRequest{Request: request}
		input, err := ParseRequestData(httpRequest)
		if err != nil {
			return nil, err
		}

		err = handler.Post(input)
		if err != nil {
			return nil, err
		}

		return "you have successfully added a new user", nil
	}
}

// HandlePut is a higher-order function that takes an IHandler as input and returns a function that handles a PUT request.
func HandlePut(handler IHandler) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(response http.ResponseWriter, request *http.Request) (any, error) {
		id := request.PathValue("id")

		httpRequest := &HttpRequest{Request: request}
		input, err := ParseRequestData(httpRequest)
		if err != nil {
			return nil, err
		}

		err = handler.Put(id, input)
		if err != nil {
			return nil, err
		}

		return "You have successfully updated an activity with id: " + id, err
	}
}

// HandleDelete is a higher-order function that takes an IHandler as input and returns a function that handles a DELETE request.
func HandleDelete(handler IHandler) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(response http.ResponseWriter, request *http.Request) (any, error) {
		id := request.PathValue("id")

		err := handler.Delete(id)
		if err != nil {
			return nil, err
		}

		return "you have successfully deleted user with id: " + id, err
	}
}
