//go:build online

package main

import (
	"backend/server"
	"fmt"
)

func main() {
	apiServer := server.New("lmbek.dk", "4040", api.AddRoutes())
	err := apiServer.Start()

	fmt.Println(err)
}
