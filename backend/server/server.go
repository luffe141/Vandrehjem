package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr string
	port string
	mux  *http.ServeMux
}

func New(addr string, port string, mux *http.ServeMux) *Server {
	return &Server{addr, port, mux}
}

func (server *Server) Start() error {
	// start server (listen on port)
	fmt.Println("Starting api server on: http://" + server.addr + ":" + server.port)
	fmt.Println("API like activities is on: http://" + server.addr + ":" + server.port + "/api/activities/")
	fmt.Println("ADMIN on: http://" + server.addr + ":" + server.port + "/admin/")
	return server.listen()
}
