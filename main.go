package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/bgiegel/todorest/model"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/tasks", Tasks)
	router.HandleFunc("/tasks/{taskId}", Task)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index handle http get request to root element
func Index(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Hello, %q", html.EscapeString(req.URL.Path))
}

// Tasks return all recorded tasks
func Tasks(response http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(response, "Hello, %q \n", html.EscapeString(req.URL.Path))

	tasks := model.Tasks{
		model.Task{Description: "Tache 1"},
		model.Task{Description: "Tache 2"},
	}

	json.NewEncoder(response).Encode(tasks)
}

// Task return the task corrresponding to the ID
func Task(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Hello, %q", html.EscapeString(req.URL.Path))
}
