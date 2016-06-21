package main

import (
    "net/http"
)
func main() {
    http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
        rw.Write([]byte("PING\n"))
    })

    http.ListenAndServe(":3000",nil)
}