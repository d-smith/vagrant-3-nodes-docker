package main

import (
    "net/http"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
        ping,err := http.Get("pingsvc:3000")
        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }

        pong,err := http.Get("pongsvc:4000")
        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }

        defer ping.Body.Close()
        defer pong.Body.Close()

        pingData,_ := ioutil.ReadAll(ping.Body)
        pongData,_ := ioutil.ReadAll(pong.Body)

        rw.Write(pingData)
        rw.Write(pongData)
    })

    http.ListenAndServe(":8080",nil)
}