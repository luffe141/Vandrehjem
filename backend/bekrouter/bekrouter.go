package bekrouter

import (
	"backend/api/controllers"
	"github.com/lmbek/bekrouter/wrapper"
	"net/http"
	"path"
)

// Router is a type that represents a Router for handling HTTP requests.
type Router struct {
	Routes []Route
	Mux    *http.ServeMux
}

// New creates a new instance of the router.
// It initializes a router struct with an empty route list and a new http.ServeMux.
// Returns a pointer to the newly created router instance.
//
// Example:
//
//	myRouter := New()
func New(routes ...Route) *Router {
	return &Router{
		Routes: routes,
		Mux:    http.NewServeMux(),
	}
}

// Route represents a route for handling HTTP requests.
// It contains the pattern for matching a request URL and the handler function for processing the request.
type Route struct {
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// AddRoute appends a new Route to the list of routes in the router.
// It takes a pattern string and a handler function as parameters.
// The pattern represents the URL pattern to match, and the handler is the function to be executed when the pattern is matched.
// The method creates a new Route struct with the given pattern and handler, and appends it to the Routes list.
// It also adds the pattern and handler to the http.ServeMux of the router.
//
// Example:
//
//	myRouter.AddRoute("GET /api/users/", HandlerFunc)
//	myRouter.AddRoute("GET /api/users/{id}/", HandlerFunc)
//	myRouter.AddRoute("POST /api/users/", HandlerFunc)
//	myRouter.AddRoute("PUT /api/users/{id}/", HandlerFunc)
//	myRouter.AddRoute("DELETE /api/users/{id}/", HandlerFunc)
//
// Parameters:
// - pattern (string): The URL pattern to match.
// - handler (http.HandlerFunc): The function to be executed when the pattern is matched.
func (router *Router) AddRoute(pattern string, handler http.HandlerFunc) {
	router.Routes = append(router.Routes, Route{
		Pattern:     pattern,
		HandlerFunc: handler,
	})
	router.Mux.HandleFunc(pattern, handler)
}

// AddRestfulJSONRoute adds the restful routes to the router for JSON responses.
// It takes a route string and an IHandler interface as parameters.
// The route parameter represents the base route for the endpoints.
// The handler parameter is an implementation of the IHandler interface, which defines the methods GetALl, GetById, Post, Put, and Delete.
// The method constructs the URL patterns using the route parameter and adds them to the router using the AddRoute method.
// The endpoint URLs are constructed as follows:
//
//	GET /api/{route}/          for retrieval of all entities
//	GET /api/{route}/{id}/     for retrieval of a specific entity by ID
//	POST /api/{route}/         for creating a new entity
//	PUT /api/{route}/{id}/     for updating a specific entity by ID
//	DELETE /api/{route}/{id}/  for deleting a specific entity by ID
//
// Parameter:
// - route (string): The base route for the endpoints.
// - handler (controllers.IHandler): An implementation of the IHandler interface.
//
// Example:
//
//	myRouter.AddRestfulJSONRoute("/api/activities/", models.Activity{})
//
// Note: The IHandler interface must be implemented by a struct and should define the methods GetAll, GetById, Post, Put, and Delete.
//
//	The wrapper.Json function wraps the IHandler methods with JSON response generation.
//	The router.AddRoute method is used to add the route patterns to the router.
//	The models.Activity struct is an example implementation of the IHandler interface.
func (router *Router) AddRestfulJSONRoute(route string, handler controllers.IHandler) {
	router.AddRoute("GET "+path.Join(route)+"/", wrapper.Json(controllers.HandleGet(handler)))            // GET /api/<your route here>/
	router.AddRoute("GET "+path.Join(route)+"/{id}/", wrapper.Json(controllers.HandleGetById(handler)))   // GET /api/<your route here>/{id}/
	router.AddRoute("POST "+path.Join(route)+"/", wrapper.Json(controllers.HandlePost(handler)))          // POST /api/<your route here>/
	router.AddRoute("PUT "+path.Join(route)+"/{id}/", wrapper.Json(controllers.HandlePut(handler)))       // PUT /api/<your route here>/{id}/
	router.AddRoute("DELETE "+path.Join(route)+"/{id}/", wrapper.Json(controllers.HandleDelete(handler))) // DELETE /api/<your route here>/{id}/
}
