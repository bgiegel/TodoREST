package web

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bgiegel/TodoREST/model"
	"github.com/gorilla/mux"
)

// Index handle http get request to root element
func (app *TodoApp) Index(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Hello, %q", html.EscapeString(req.URL.Path))
}

// Tasks return all recorded tasks
func (app *TodoApp) Tasks(response http.ResponseWriter, req *http.Request) {
	tasks := app.TaskRepo.AllTasks()

	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(tasks); err != nil {
		panic(err)
	}
}

// Task return the task corrresponding to the ID
func (app *TodoApp) Task(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskID, _ := strconv.Atoi(vars["taskID"])

	task := app.TaskRepo.FindTask(taskID)

	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(task); err != nil {
		panic(err)
	}
}

// TaskCreate create a new Task
func (app *TodoApp) TaskCreate(response http.ResponseWriter, req *http.Request) {
	var task model.Task
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &task); err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(response).Encode(err); err != nil {
			panic(err)
		}
	}

	task.ID = app.TaskRepo.CreateTask(task)
	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(response).Encode(task); err != nil {
		panic(err)
	}
}

//TaskDelete delete a task with corresponding taskID
func (app *TodoApp) TaskDelete(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskID, _ := strconv.Atoi(vars["taskID"])

	app.TaskRepo.DestroyTask(taskID)

	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusOK)
}

// TaskCreate create a new Task
func (app *TodoApp) TaskUpdate(response http.ResponseWriter, req *http.Request) {
	var task model.Task
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err = req.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &task); err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(response).Encode(err)
		panic(err)
	}

	if err = app.TaskRepo.UpdateTask(task); err != nil {
		log.Panic(err)
	} else {
		log.Printf("Task %d updated with description : '%v' \n", task.ID, task.Description)
	}

	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(response).Encode(task); err != nil {
		panic(err)
	}
}
