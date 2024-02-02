package server

import (
	"backend/api"
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
	// set up routing
	mux := api.AddRoutes()
	// start server (listen on port)
	fmt.Println("Starting api server on: http://" + server.addr + ":" + server.port)
	return server.Listen(mux)
}
