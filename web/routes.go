package web

import (
	"net/http"
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

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"Tasks", "GET", "/tasks", Tasks},
	Route{"Task", "GET", "/tasks/{taskID}", Task},
	Route{"CreateTask", "POST", "/tasks", TaskCreate},
	Route{"DeleteTask", "DELETE", "/tasks/{taskID}", TaskDelete},
	Route{"UpdateTask", "PUT", "/tasks", TaskUpdate},
}