package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
}
