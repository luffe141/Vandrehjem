package handlers

import (
	"fmt"
	"net/http"
)

type ICrudable interface {
	GetAll() (any, error)
	GetById(id string) (any, error)
	Post(data map[string]any) error
	Put(id string, data map[string]any) error
	Delete(id string) error
}

func HandleGet(crudable ICrudable) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(w http.ResponseWriter, req *http.Request) (any, error) {
		fmt.Println("1")
		return crudable.GetAll()
	}
}

func HandleGetById(crudable ICrudable) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(response http.ResponseWriter, request *http.Request) (any, error) {
		id := request.PathValue("id")
		return crudable.GetById(id)
	}
}

func HandlePost(crudable ICrudable) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(response http.ResponseWriter, request *http.Request) (any, error) {
		data, err := parseRequestData(request)
		if err != nil {
			return nil, err
		}

		err = crudable.Post(data)
		if err != nil {
			return nil, err
		}

		return "you have successfully added a new user", nil
	}
}

func HandlePut(crudable ICrudable) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(response http.ResponseWriter, request *http.Request) (any, error) {
		id := request.PathValue("id")

		// Parse request data and map it to Activities struct
		data, err := parseRequestData(request)
		if err != nil {
			return nil, err
		}

		err = crudable.Put(id, data)
		if err != nil {
			return nil, err
		}

		return "You have successfully updated an activity with id: " + id, err
	}
}

func HandleDelete(crudable ICrudable) func(http.ResponseWriter, *http.Request) (any, error) {
	return func(response http.ResponseWriter, request *http.Request) (any, error) {
		id := request.PathValue("id")

		err := crudable.Delete(id)
		if err != nil {
			return nil, err
		}

		return "you have successfully deleted user with id: " + id, err
	}
}
