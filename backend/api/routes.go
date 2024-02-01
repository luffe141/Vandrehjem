package api

import (
	"backend/api/handlers"
	"backend/api/users"
	"net/http"
)

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// routes
	mux = apiBaseRoutes(mux)
	mux.HandleFunc("GET /api/activities/", handlers.HandleGetActivities)
	mux.HandleFunc("GET /api/activities/{id}/", handlers.HandleGetByIdActivities)
	mux.HandleFunc("POST /api/activities/", handlers.HandlePostActivities)
	mux.HandleFunc("PUT /api/activities/{id}/", handlers.HandlePutActivities)
	mux.HandleFunc("DELETE /api/activities/{id}/", handlers.HandleDeleteActivities)

	mux = exampleRoutes(mux)
	return mux
}

func apiBaseRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /api/", RouteDontExist)
	mux.HandleFunc("POST /api/", RouteDontExist)
	mux.HandleFunc("PUT /api/", RouteDontExist)
	mux.HandleFunc("DELETE /api/", RouteDontExist)
	return mux
}

func exampleRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /api/users/", users.GetAll)
	mux.HandleFunc("GET /api/users/{id}/", users.Get)
	mux.HandleFunc("POST /api/users/", users.Post)
	mux.HandleFunc("PUT /api/users/{id}/", users.Put)
	mux.HandleFunc("DELETE /api/users/{id}/", users.Delete)

	mux.HandleFunc("GET /api/users", RouteDontExistMissingSlash)
	mux.HandleFunc("POST /api/users", RouteDontExistMissingSlash)
	mux.HandleFunc("PUT /api/users", RouteDontExistMissingSlash)
	mux.HandleFunc("DELETE /api/users", RouteDontExistMissingSlash)
	mux.HandleFunc("PUT /api/users/", RouteMalformed)
	mux.HandleFunc("DELETE /api/users/", RouteMalformed)
	return mux
}
