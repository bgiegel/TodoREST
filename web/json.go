package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/bgiegel/TodoREST/model"
)

func ReadTaskFromRequest(response http.ResponseWriter, req *http.Request) (task model.Task) {
	body := extractBody(req)

	if err := json.Unmarshal(body, &task); err != nil {
		unprocessableEntityResponse(response, err)
	}

	return
}

func unprocessableEntityResponse(response http.ResponseWriter, err error) {
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(422) // unprocessable entity
	if err := json.NewEncoder(response).Encode(err); err != nil {
		panic(err)
	}
}

func extractBody(req *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	return body
}

func RespondWithTask(response http.ResponseWriter, task model.Task) {
	jsonOKResponse(response, task)
}

func RespondWithTasks(response http.ResponseWriter, tasks []model.Task) {
	jsonOKResponse(response, tasks)
}

func jsonOKResponse(response http.ResponseWriter, value interface{}) {
	response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(response).Encode(value); err != nil {
		panic(err)
	}
}
