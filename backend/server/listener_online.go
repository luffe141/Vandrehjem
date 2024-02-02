//go:build online

package server

import "net/http"

// Listen starts the HTTP server and listens on the specified address and port.
func (server webserver) Listen() error {
	return http.ListenAndServeTLS(server.addr+":"+server.port, "path to certificate file", mux)
}
