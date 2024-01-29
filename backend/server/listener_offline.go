//go:build !online

package server

import "net/http"

func (server webserver) Listen() error {
	return http.ListenAndServe(server.addr+":"+server.port, nil)
}
