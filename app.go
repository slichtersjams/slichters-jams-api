package app

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
    "encoding/json"
    "./jamstore"
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
    if r.Body == nil {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }

    decoder := json.NewDecoder(r.Body)
    var jam jamstore.Jam

    err := decoder.Decode(&jam)

    if err != nil {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    }
}

func getResponse(responseIndex int) string  {
    if responseIndex == 1 {
        return "Not a Jam!"
    }
    return "Jam!"
}