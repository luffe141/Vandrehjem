package server_side_render

import (
	"html/template"
	"net/http"
)

func ServerSideRender(response http.ResponseWriter, file any, htmlTemplate *template.Template) error {
	// Execute the template, passing any data if needed
	err := htmlTemplate.Execute(response, file)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return err
}
