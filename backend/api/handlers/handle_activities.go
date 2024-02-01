package handlers

import (
	"backend/api/controllers"
	"fmt"
	"net/http"
	"strings"
)

type Activity struct {
	_Id  string
	Name string
	Age  int
	Img  string
}

func mapToActivities(dataMap map[string]any) (*Activity, error) {
	var activity Activity

	for key, value := range dataMap {
		switch strings.ToLower(key) {
		case "name":
			activity.Name, _ = value.(string)
		case "age":
			if activity.Age == 0 {
				activity.Age = -1
			} else {
				activity.Age, _ = value.(int)
			}

		case "img":
			activity.Img, _ = value.(string)
		}
	}

	return &activity, nil
}

func HandleGetActivities(response http.ResponseWriter, request *http.Request) {
	data, err := controllers.GetActivities()
	if err != nil {
		fmt.Println(err)
	}
	_, _ = fmt.Fprint(response, data)
}

func HandleGetByIdActivities(response http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	data, err := controllers.GetActivity(id)

	if err != nil {
		fmt.Println(err)
	}

	_, _ = fmt.Fprint(response, data)
}

func HandlePostActivities(response http.ResponseWriter, request *http.Request) {
	dataMap, err := parseRequestData(request)
	if err != nil {
		http.Error(response, "Error parsing request: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = fmt.Fprint(response, dataMap)
	_, _ = fmt.Fprint(response, dataMap)

	data, err := mapToActivities(dataMap)
	if err != nil {
		http.Error(response, "Error mapping request data to Activities: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = controllers.PostActivity(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	_, _ = fmt.Fprint(response, "you have successfully added a new user")

}

func HandlePutActivities(response http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")

	// Parse request data and map it to Activities struct
	dataMap, err := parseRequestData(request)
	if err != nil {
		_, _ = fmt.Fprint(response, "Error parsing request data: "+err.Error())
		return
	}

	data, err := mapToActivities(dataMap)
	if err != nil {
		_, _ = fmt.Fprint(response, "Error mapping request data to Activities: "+err.Error())
		return
	}

	err = controllers.PutActivity(id, data)
	if err != nil {
		_, _ = fmt.Fprint(response, err)
		return
	}

	_, _ = fmt.Fprint(response, "You have successfully updated an activity.")

}

func HandleDeleteActivities(response http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")

	err := controllers.DeleteActivity(id)
	if err != nil {
		_, _ = fmt.Fprint(response, err)
		return
	}

	_, _ = fmt.Fprint(response, "you have succesfully deleted user with id: ", id)

}
