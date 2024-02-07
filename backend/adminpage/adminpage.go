package adminpage

import (
	serverSideRender "backend/adminpage/server-side-render"
	"backend/api/database/mongodb"
	"embed"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"net/http"
)

//go:embed frontend/index.html
var content embed.FS

func ServeAdminPage(response http.ResponseWriter, request *http.Request) {
	// Retrieve the template file from the embedded FS
	htmlTemplate, err := template.ParseFS(content, "frontend/index.html")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	MongodbConnection := "mongodb://localhost:27017"
	DatabaseName := "vandrerhjem"
	CollectionName := "rooms"

	storage := mongodb.NewStorage(MongodbConnection, DatabaseName, CollectionName)

	rooms, err := storage.ReadData(bson.M{})
	if err != nil {
		http.Error(response, err.Error(), 500)
	}

	err = serverSideRender.ServerSideRender(response, rooms, htmlTemplate)
	if err != nil {
		http.Error(response, err.Error(), 500)
	}

}
