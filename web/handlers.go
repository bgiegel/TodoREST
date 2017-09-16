package web

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index handle http get request to root element
func (app *TodoApp) Index(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Hello, %q", html.EscapeString(req.URL.Path))
}

// Tasks return all recorded tasks
func (app *TodoApp) Tasks(response http.ResponseWriter, req *http.Request) {
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
