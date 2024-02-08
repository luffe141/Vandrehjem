//go:build production

package adminpage

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed frontend/index.html
var content embed.FS

var patterns = "frontend/index.html"

func GetHTMLTemplate(response http.ResponseWriter, request *http.Request) *template.Template {
	// Retrieve the template file from the embedded FS
	htmlTemplate, err := template.ParseFS(content, patterns)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return htmlTemplate
}
