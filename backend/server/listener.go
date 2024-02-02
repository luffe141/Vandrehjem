//go:build !online

package server

import "net/http"

func (server webserver) Listen(mux *http.ServeMux) error {
	return http.ListenAndServe(server.addr+":"+server.port, mux)
}
