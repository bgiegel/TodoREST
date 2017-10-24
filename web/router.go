package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
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
	var protectedRoutes = Routes{
		Route{"Index", "GET", "/", router.App.Index},
		Route{"Tasks", "GET", "/tasks", router.App.Tasks},
		Route{"Task", "GET", "/tasks/{taskID}", router.App.Task},
		Route{"CreateTask", "POST", "/tasks", router.App.TaskCreate},
		Route{"DeleteTask", "DELETE", "/tasks/{taskID}", router.App.TaskDelete},
		Route{"UpdateTask", "PUT", "/tasks", router.App.TaskUpdate},
	}

	for _, route := range protectedRoutes {
		handler := validate(Logger(route.HandlerFunc, route.Name))
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	tokenRoute := Route{"SetToken", "GET", "/settoken", router.App.SetToken}
	router.Methods(tokenRoute.Method).Path(tokenRoute.Pattern).Name(tokenRoute.Name).Handler(Logger(tokenRoute.HandlerFunc, tokenRoute.Name))

	return router
}

func validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// If no Auth cookie is set then return a 404 not found
		cookie, err := req.Cookie("Auth")
		if err != nil {
			http.NotFound(res, req)
			return
		}

		// Return a Token using the cookie
		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error){
			// Make sure token's signature wasn't changed
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			http.NotFound(res, req)
			return
		}

		// Grab the tokens claims and pass it into the original request
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), "jwt-token", *claims)
			next.ServeHTTP(res, req.WithContext(ctx))
		} else {
			http.NotFound(res, req)
			return
		}
	})
}
