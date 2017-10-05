package app

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
)

func init() {
    http.HandleFunc("/", getHandler)
    http.HandleFunc("/jams", jamPostHandler)
    rand.Seed(time.Now().UTC().UnixNano())
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    fmt.Fprint(w, getResponse(rand.Intn(2)))
}

func jamPostHandler(w http.ResponseWriter, r *http.Request) {

}

func getResponse(responseIndex int) string  {
    if responseIndex == 1 {
        return "Not a Jam!"
    }
    return "Jam!"
}