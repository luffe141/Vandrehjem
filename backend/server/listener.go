//go:build !online

package server

import "net/http"

func (server *Server) listen() error {
	return http.ListenAndServe(server.addr+":"+server.port, server.mux)
}
