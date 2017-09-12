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
	"github.com/bgiegel/TodoREST/repo"
	"github.com/gorilla/mux"
)

// Index handle http get request to root element
func Index(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Hello, %q", html.EscapeString(req.URL.Path))
}

// Tasks return all recorded tasks
func Tasks(response http.ResponseWriter, req *http.Request) {
	tasks := repo.RepoAllTasks()

	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(tasks); err != nil {
		panic(err)
	}
}

// Task return the task corrresponding to the ID
func Task(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskID, _ := strconv.Atoi(vars["taskID"])

	task := repo.RepoFindTask(taskID)
	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(task); err != nil {
		panic(err)
	}
}

// TaskCreate create a new Task
func TaskCreate(response http.ResponseWriter, req *http.Request) {
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
	log.Println("Creation d'une nouvelle tache !")
	log.Println("Description : " + task.Description)

	task = repo.RepoCreateTask(task)
	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(task); err != nil {
		panic(err)
	}
}
