package adminpage

import (
	"backend/adminpage/data"
	serverSideRender "backend/adminpage/server-side-render"
	"fmt"
	"html/template"
	"net/http"
)

func ServeAdminPage(response http.ResponseWriter, request *http.Request) {
	htmlTemplate, err := GetHTMLTemplate(response, request)

	if err != nil {
		http.Error(response, err.Error(), 500)
	}

	// Fetch data
	galleries, err := data.GetGalleries()
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}
	fmt.Println("galleries")
	fmt.Println(galleries)
	/*
		rooms, err := data.GetRooms()
		if err != nil {
			http.Error(response, err.Error(), 500)
			return
		}
		fmt.Println("rooms")
		fmt.Println(rooms)
	*/
	// Prepare gallery
	templateData, err := serverSideRender.Prepare(galleries, htmlTemplate)
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}

	// Combine prepared data and render all
	finalTemplate := template.New("final")
	finalTemplate, err = finalTemplate.Parse(templateData)
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}

	// Execute the final template with any required data (if the combined template needs it)
	err = finalTemplate.Execute(response, nil)
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}

}
