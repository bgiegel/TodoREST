package main

import (
	"log"
	"net/http"

	"github.com/bgiegel/TodoREST/repo"
	"github.com/bgiegel/TodoREST/web"

	// pq est le driver pour se connecter à la base PostgreSQL
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting TodoRest on Port :8080")

	app := &web.TodoApp{TaskRepo: repo.GetTaskRepository()}

	router := web.NewRouter(app).InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
