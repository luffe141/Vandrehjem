package api

import (
	"backend/api/controllers/mongodb"
	"backend/api/handlers"
	"backend/api/users"
	"net/http"
)

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// routes
	mux = apiBaseRoutes(mux)

	activity := &mongodb.Activity{}
	aboutUs := &mongodb.AboutUs{}

	//mux.HandleFunc("/api/activities/", HandleGet(activity))
	//mux.HandleFunc("/api/sliders/", HandleGet(slider))

	mux.HandleFunc("GET /api/activities", JsonWrapper(handlers.HandleGet(activity)))
	mux.HandleFunc("GET /api/activities/{id}/", JsonWrapper(handlers.HandleGetById(activity)))
	mux.HandleFunc("POST /api/activities/", JsonWrapper(handlers.HandlePost(activity)))
	mux.HandleFunc("PUT /api/activities/{id}/", JsonWrapper(handlers.HandlePut(activity)))
	mux.HandleFunc("DELETE /api/activities/{id}/", JsonWrapper(handlers.HandleDelete(activity)))

	mux.HandleFunc("GET /api/aboutus", JsonWrapper(handlers.HandleGet(aboutUs)))
	mux.HandleFunc("GET /api/aboutus/{id}/", JsonWrapper(handlers.HandleGetById(aboutUs)))
	mux.HandleFunc("POST /api/aboutus/", JsonWrapper(handlers.HandlePost(aboutUs)))
	mux.HandleFunc("PUT /api/aboutus/{id}/", JsonWrapper(handlers.HandlePut(aboutUs)))
	mux.HandleFunc("DELETE /api/aboutus/{id}/", JsonWrapper(handlers.HandleDelete(aboutUs)))

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
