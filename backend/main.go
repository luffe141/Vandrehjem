//go:build !online

package main

import (
	"backend/api"
	"backend/server"
	"fmt"
)

func main() {
	apiServer := server.New("127.0.0.1", "8070", api.AddRoutes())
	err := apiServer.Start()

	fmt.Println(err)
}
