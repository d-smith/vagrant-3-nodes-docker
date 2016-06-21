package main

import (
    "net/http"
)
func main() {
    http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
        rw.Write([]byte("PONG\n"))
    })

    http.ListenAndServe(":4000",nil)
}