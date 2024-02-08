//go:build !embedded && !production

package adminpage

import (
	"net/http"
	"os"
)

func GetHTMLTemplate(response http.ResponseWriter, request *http.Request) (string, error) {
	filePath := "./adminpage/frontend/admin.gohtml"
	// Check if the file exists before attempting to open
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		http.Error(response, "File does not exist", http.StatusInternalServerError)
	}
	// Parse the HTML file directly into a template
	file, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	return string(file), err
}
