package app

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Jam!")
}

func getResponse(responseIndex int) string  {
    return "Jam!"
}