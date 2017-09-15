package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TodoRouter struct {
	*mux.Router
	App *TodoApp
}

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
func NewRouter(app *TodoApp) *TodoRouter {
	router := mux.NewRouter().StrictSlash(true)
	todoRouter := &TodoRouter{router, app}

	return todoRouter
}

// InitRoutes initialise la liste des routes d'une application
func (router *TodoRouter) InitRoutes() *TodoRouter {
	var routes = Routes{
		Route{"Index", "GET", "/", router.App.Index},
		Route{"Tasks", "GET", "/tasks", router.App.Tasks},
		Route{"Task", "GET", "/tasks/{taskID}", router.App.Task},
		Route{"CreateTask", "POST", "/tasks", router.App.TaskCreate},
		Route{"DeleteTask", "DELETE", "/tasks/{taskID}", router.App.TaskDelete},
		Route{"UpdateTask", "PUT", "/tasks", router.App.TaskUpdate},
	}

	for _, route := range routes {
		handler := Logger(route.HandlerFunc, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}
