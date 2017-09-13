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
	repo.InitDB()
	router := web.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
