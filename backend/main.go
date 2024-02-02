//go:build !online

package main

import (
	"backend/server"
	"fmt"
)

func main() {
	apiServer := server.New("localhost", "8070")
	err := apiServer.Start()

	fmt.Println(err)
}
