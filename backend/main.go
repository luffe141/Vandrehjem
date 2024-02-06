//go:build !online

package main

import (
	"backend/api/mvc/views"
	"backend/server"
	"fmt"
)

func main() {
	apiServer := server.New("127.0.0.1", "8070", views.AddRoutes())
	err := apiServer.Start()

	fmt.Println(err)
}
