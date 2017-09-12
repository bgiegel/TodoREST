package main

import (
	"log"
	"net/http"

	"github.com/bgiegel/TodoREST/web"
)

func main() {
	router := web.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
