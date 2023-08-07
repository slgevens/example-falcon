package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello, my name is Falcon")
}

func main() {
    http.HandleFunc("/", hello)
    port := ":4000"
    fmt.Println("Server running at http://localhost" + port)
    http.ListenAndServe(port, nil)
}
