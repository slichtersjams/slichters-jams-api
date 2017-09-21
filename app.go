package app

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
)

func init() {
    http.HandleFunc("/", handler)
    rand.Seed(time.Now().UTC().UnixNano())
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    fmt.Fprint(w, getResponse(rand.Intn(2)))
}

func getResponse(responseIndex int) string  {
    if responseIndex == 1 {
        return "Not a Jam!"
    }
    return "Jam!"
}