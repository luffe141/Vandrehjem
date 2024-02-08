package serverSideRender

import (
	"bytes"
	"html/template"
)

// Prepare compiles the template with given data and returns the result as a string.
func Prepare(data any, html string) (string, error) {
	var buffer bytes.Buffer

	// Parse the HTML template string
	htmlTemplate, err := template.New("template").Parse(html)
	if err != nil {
		return "", err
	}
	// Execute the template into a buffer instead of directly to response
	err = htmlTemplate.Execute(&buffer, data)
	if err != nil {
		return "", err
	}

	// Return the buffer content as a string
	return buffer.String(), nil
}
