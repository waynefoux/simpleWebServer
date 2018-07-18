package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Welcome!")
}

func TodoIndex(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "TODO")
}

func TodoShow(writer http.ResponseWriter, request *http.Request) {
    vars := mux.Vars(request)
    todoId := vars["todoId"]
    fmt.Fprintf(writer, "Toddo show:", todoId)
}
