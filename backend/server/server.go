package server

import (
	"fmt"
)

type webserver struct {
	addr string
	port string
}

func New(addr string, port string) *webserver {
	return &webserver{addr, port}
}

func (server webserver) Start() error {
	fmt.Println("Starting api server on: " + server.addr + ":" + server.port)
	return server.Listen()
}
