package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Helper function to parse the form data and the query parameters
func parseFormAndQuery(request *http.Request) (map[string]string, error) {
	err := request.ParseForm()
	if err != nil {
		return nil, err
	}

	flatForm := make(map[string]string)
	for key, values := range request.Form {
		if len(values) > 0 {
			flatForm[key] = values[0]
		}
	}
	return flatForm, nil
}

// Helper function to parse the JSON body
func parseJSONBody(request *http.Request) (map[string]any, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Error in reading the request body: ", err)
		return nil, err
	}

	var data map[string]any
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error in unmarshalling the JSON: ", err)
		return nil, err
	}

	return data, nil
}

func parseRequestData(request *http.Request) (map[string]any, error) {
	data := make(map[string]any)

	if request.URL.RawQuery != "" {
		// Parse the form data.
		formData, err := parseFormAndQuery(request)
		if err != nil {
			return nil, err
		}

		for k, v := range formData {
			data[k] = v
		}
	}

	if request.ContentLength != 0 {
		// Parse the JSON body if the request contain one.
		jsonData, err := parseJSONBody(request)
		if err != nil {
			return nil, err
		}

		for k, v := range jsonData {
			data[k] = v
		}
	}

	return data, nil
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
