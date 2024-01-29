//go:build online

package server

import "net/http"

func (server webserver) Listen() error {
	return http.ListenAndServeTLS(server.addr+":"+server.port, "path to certificate file")
}
