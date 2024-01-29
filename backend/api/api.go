package api

import "net/http"

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("GET /users/", handleGetUsers)
	mux.HandleFunc("GET /users/{id}/", handleGetUsersId)
	mux.HandleFunc("GET /usersa/", handleGetUsersa)
	mux.HandleFunc("GET /usersb/", handleGetUsersb)

	return mux
}
