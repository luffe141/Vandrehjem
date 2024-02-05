package controllers

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ParsedBody map[string]any

type HttpRequest struct {
	Request *http.Request
}

func ParseFormAndQuery(request *HttpRequest) (map[string]any, error) {
	err := request.Request.ParseForm()
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	for name, values := range request.Request.Form {
		data[name] = values[0]
	}

	return data, nil
}

func ParseJSONBody(request *HttpRequest) (map[string]any, error) {
	defer request.Request.Body.Close()

	data := make(map[string]any)
	err := json.NewDecoder(request.Request.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*

// todo: feature request: generic XML parsing
type Note struct {
	To string `xml:"to"`
}

// ParseXMLBody parses the XML body of an HttpRequest object.
// It closes the request body using defer to ensure proper resource cleanup.
// The function creates an empty map, data, of type map[string]any to hold the parsed XML data.
// It then uses xml.NewDecoder to create a new XML decoder from the request body and calls the Decode method to decode the XML into the data map.
// If an error occurs during the decoding process, the function returns nil and the error.
// If decoding is successful, it returns the data map and a nil error.
func ParseXMLBody(request *HttpRequest) (*Note, error) {
	defer request.Request.Body.Close()
	data := new(Note)
	err := xml.NewDecoder(request.Request.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
*/

// ParseHTMLBody parses the HTML body from an HttpRequest object.
// It closes the request body using the `defer` statement.
// The parsed HTML body is stored in a map[string]any, with the key "html" and the raw request body as the value.
// The function does not return an error, assuming the parsing of the HTML body will always be successful.
// The `ParseHTMLBody` function is used in the `ParseRequestData` function to handle requests with a Content-Type of "text/html".
func ParseHTMLBody(request *HttpRequest) (map[string]any, error) {
	data := make(map[string]any)

	_, err := io.ReadAll(request.Request.Body)
	if err != nil {
		return nil, err
	}

	request.Request.Body = http.NoBody

	data["html"] = request.Request.Body

	return data, nil
}

func ParseCSVBody(request *HttpRequest) ([]map[string]any, error) {
	defer request.Request.Body.Close()

	var data []map[string]any
	reader := csv.NewReader(request.Request.Body)

	var header []string
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV data: %v\n", err)
		return nil, err
	}

	fmt.Printf("Number of lines in CSV: %d\n", len(records))

	for idx, record := range records {
		if idx == 0 {
			header = record
			continue
		}

		row := make(map[string]any)
		for i, cell := range record {
			colname := header[i]
			row[colname] = cell
		}
		data = append(data, row)
	}

	return data, nil
}

// ParseRequestData parses the request data based on the Content-Type header.
// It takes in an HttpRequest object and returns a map[string]any with the parsed data
// and an error if parsing fails.
// If the Content-Type header is empty, it returns the data map immediately.
// If the Content-Type header specifies "application/json", it calls the ParseJSONBody function to parse the JSON body.
// If the Content-Type header specifies "text/csv", it calls the ParseCSVBody function to parse the CSV body.
// If the Content-Type header specifies "text/html", it calls the ParseHTMLBody function to parse the HTML body.
// If the Content-Type header specifies "application/x-www-form-urlencoded", it assumes the form data has already been parsed.
// If the Content-Type header is unrecognized, it returns an error indicating unsupported content type.
// The function merges the parsed data into the data map using the key-value pairs from the parsed data map.
// The function returns the merged data map and nil error if successful.
func ParseRequestData(request *HttpRequest) (any, error) {
	data := make(map[string]any)

	contentType := request.Request.Header.Get("Content-Type")

	// Parse form and url parameters
	formData, err := ParseFormAndQuery(request)
	if err != nil {
		return nil, err
	}
	// Merge parsed form data into data map
	for k, v := range formData {
		data[k] = v
	}

	// if no Content-Type header, end here
	if contentType == "" {
		return data, nil
	}

	switch {
	case strings.Contains(contentType, "application/json"):
		bodyData, err := ParseJSONBody(request)
		if err != nil {
			return nil, err
		}
		// Merge parsed body data
		for k, v := range bodyData {
			data[k] = v
		}
	case strings.Contains(contentType, "text/csv"):
		records, err := ParseCSVBody(request)
		if err != nil {
			return nil, err
		}
		data["records"] = records
	case strings.Contains(contentType, "text/html"):
		bodyData, err := ParseHTMLBody(request)
		if err != nil {
			return nil, err
		}
		// Merge parsed body data
		for k, v := range bodyData {
			data[k] = v
		}
	case strings.Contains(contentType, "application/x-www-form-urlencoded"):
		//already parsed above as form data
	default:
		return nil, errors.New("unsupported content type")
	}

	return data, nil
}
