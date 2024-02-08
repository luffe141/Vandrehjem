//go:build !online
// +build !online

package main

import (
	"backend/adminpage"
	"backend/api/mvc/views"
	"backend/server"
	"fmt"
)

func main() {
	routes := views.AddRoutes()
	routes.HandleFunc("/admin/", adminpage.ServeAdminPage)
	apiServer := server.New("127.0.0.1", "8070", routes)
	err := apiServer.Start()
	if err != nil {
		fmt.Println(err)
	}

}
