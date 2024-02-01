package handlers

import (
	"backend/api/controllers"
	"net/http"
)

func HandleGetData(response http.ResponseWriter, request *http.Request) (any, error) {
	return controllers.GetActivities()
}

func HandleGetDataById(response http.ResponseWriter, request *http.Request) (any, error) {
	id := request.PathValue("id")
	return controllers.GetActivity(id)
}

func HandlePostData(response http.ResponseWriter, request *http.Request) (any, error) {
	dataMap, err := parseRequestData(request)
	if err != nil {
		return nil, err
	}

	data, err := mapToActivities(dataMap)
	if err != nil {
		return nil, err
	}

	err = controllers.PostActivity(data)
	if err != nil {
		return nil, err
	}

	return "you have successfully added a new user", nil
}

func HandlePutData(response http.ResponseWriter, request *http.Request) (any, error) {
	id := request.PathValue("id")

	// Parse request data and map it to Activities struct
	dataMap, err := parseRequestData(request)
	if err != nil {
		return nil, err
	}

	data, err := mapToActivities(dataMap)
	if err != nil {
		return nil, err
	}

	err = controllers.PutActivity(id, data)
	if err != nil {
		return nil, err
	}

	return "You have successfully updated an activity with id: " + id, err
}

func HandleDeleteData(response http.ResponseWriter, request *http.Request) (any, error) {
	id := request.PathValue("id")

	err := controllers.DeleteActivity(id)
	if err != nil {
		return nil, err
	}

	return "you have successfully deleted user with id: " + id, err
}
