package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Helper function to parse the form data and the query parameters
func parseFormAndQuery(request *http.Request) (map[string]interface{}, error) {
	err := request.ParseForm()
	if err != nil {
		return nil, err
	}

	flatForm := make(map[string]interface{})
	for key, values := range request.Form {
		if len(values) > 0 {
			flatForm[key] = values[0]
		}
	}
	return flatForm, nil
}

// Helper function to parse the JSON body
func parseJSONBody(request *http.Request) (map[string]interface{}, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Error in reading the request body: ", err)
		return nil, err
	}
	fmt.Println("read")
	fmt.Println(body)
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error in unmarshalling the JSON: ", err)
		return nil, err
	}

	fmt.Println("data")
	fmt.Println(data)

	return data, nil
}

func parseRequestData(request *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})

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
