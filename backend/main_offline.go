//go:build !online

package main

import (
	"backend/server"
	"fmt"
)

func main() {
	apiServer := server.New("127.0.0.1", "4040")
	err := apiServer.Start()

	fmt.Println(err)
}
