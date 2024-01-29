package main

import (
	"backend/server"
	"fmt"
)

func main() {
	err := server.Start()

	fmt.Println(err)
}
