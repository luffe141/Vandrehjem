package api

import "net/http"

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// routes
	//mux.HandleFunc("GET /users/", handleGetUsers)
	//mux.HandleFunc("POST /users/", handlePostUsers)
	//mux.HandleFunc("GET /users/{id}/", handleGetUsersId)

	//mux.HandleFunc("GET /usersa/", handleGetUsersa)
	//mux.HandleFunc("GET /usersb/", handleGetUsersb)

	//mux.HandleFunc("GET /users/", tasks.GetAll)
	//mux.HandleFunc("GET /users/{id}/", tasks.Get)
	mux.HandleFunc("POST /users/", handlePostUsers)
	//mux.HandleFunc("PUT /users/{id}/", tasks.Put)
	//mux.HandleFunc("DELETE /users/{id}/", tasks.Delete)

	return mux
}
