package main

import (
    "fmt"
    "log"
    "net/http"
)

func おはようございます(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "おはようございます\n")
}

func main() {
    http.HandleFunc("/", おはようございます)
    err := http.ListenAndServe("localhost:1234", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
