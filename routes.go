package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represente une route HTTP
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes represente une liste de route HTTP
type Routes []Route

// NewRouter créé les routes définies dans routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"Tasks", "GET", "/tasks", Tasks},
	Route{"Task", "GET", "/tasks/{taskId}", Task},
}
