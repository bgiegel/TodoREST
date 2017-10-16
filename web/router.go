package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"fmt"
	"github.com/auth0-community/auth0"
	"gopkg.in/square/go-jose.v2"
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
		//Route{"GetToken", "GET", "/get-token", router.App.GetToken},
	}

	for _, route := range routes {
		handler := authMiddleware(Logger(route.HandlerFunc, route.Name))
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte("w0Y0hcRiufEliTy4_1BFXMUG-_ugfVPASx0WjHRoRITqUJhYETAhqCZZyUk3r1sN")
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{"todorest"}

		configuration := auth0.NewConfiguration(secretProvider, audience, "https://todo-rest.eu.auth0.com/", jose.HS256)
		validator := auth0.NewValidator(configuration)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
