package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the Index!")
}

func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	todos := Todos{
		Todo{Name: "Write a webserver"},
		Todo{Name: "Learn Go"},
	}

	if err := json.NewEncoder(writer).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintf(writer, "Toddo show:", todoId)
}
