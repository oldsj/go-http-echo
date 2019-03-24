package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

var port = "5000"

func main() {
    log.Fatal(http.ListenAndServe(":"+port, router()))
}

func router() http.Handler {
    r := mux.NewRouter()
    r.Path("/healthz").Methods(http.MethodGet).HandlerFunc(greet)
    return r
}

func greet(w http.ResponseWriter, req *http.Request) {
    _, _ = w.Write([]byte("Hello, world! Lets try somethin else"))
}
