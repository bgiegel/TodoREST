package web

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// Initialise un Token
func (app *TodoApp) SetToken(response http.ResponseWriter, req *http.Request) {
	// Expires the token and cookie in 1 hour
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	expireCookie := time.Now().Add(time.Hour * 1)

	// We'll manually assign the claims but in production you'd insert values from a database
	claims := Claims {
		"john",
		"user",
		jwt.StandardClaims {
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000",
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := token.SignedString([]byte("secret"))

	// Place the token in the client's cookie
	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(response, &cookie)

	// Redirect the user to his profile
	http.Redirect(response, req, "/tasks", 307)
}

// Index handle http get request to root element
func (app *TodoApp) Index(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Hello, %q", html.EscapeString(req.URL.Path))
}

// Tasks return all recorded tasks
func (app *TodoApp) Tasks(response http.ResponseWriter, req *http.Request) {

	claims, ok := req.Context().Value("jwt-token").(Claims)
	if !ok{
		UnauthorizedResponse(response, fmt.Errorf("Not Authenticated"))
		return
	}
	if claims.Role != "user" {
		ForbiddenResponse(response, fmt.Errorf("Not Authorized"))
		return
	}

	log.Println(response, "Hello %s", claims.Username)

	tasks := app.TaskRepo.AllTasks()

	RespondWithTasks(response, tasks)
}

// Task return the task corrresponding to the ID
func (app *TodoApp) Task(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskID, _ := strconv.Atoi(vars["taskID"])

	task := app.TaskRepo.FindTask(taskID)

	RespondWithTask(response, task)
}

// TaskCreate create a new Task
func (app *TodoApp) TaskCreate(response http.ResponseWriter, req *http.Request) {
	task := ReadTaskFromRequest(response, req)

	task.ID = app.TaskRepo.CreateTask(task)

	RespondWithTask(response, task)
}

//TaskDelete delete a task with corresponding taskID
func (app *TodoApp) TaskDelete(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskID, _ := strconv.Atoi(vars["taskID"])

	if err := app.TaskRepo.DestroyTask(taskID); err != nil {
		log.Panic(err)
	} else {
		log.Printf("Task %d deleted \n", taskID)
	}

	response.WriteHeader(http.StatusOK)
}

// TaskUpdate create a new Task
func (app *TodoApp) TaskUpdate(response http.ResponseWriter, req *http.Request) {
	task := ReadTaskFromRequest(response, req)

	if err := app.TaskRepo.UpdateTask(task); err != nil {
		log.Panic(err)
	} else {
		log.Printf("Task %d updated with description : '%v' \n", task.ID, task.Description)
	}

	RespondWithTask(response, task)
}
