package app

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
    "encoding/json"
    "google.golang.org/appengine"
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
    var jam Jam

    err := decoder.Decode(&jam)

    if err != nil {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }

    if len(jam.JamText) == 0 {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }

    ctx := appengine.NewContext(r)

    dataStore := DataStore{ctx}

    if err = StoreJam(&dataStore, jam.JamText, jam.State); err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }
}

func getResponse(responseIndex int) string  {
    if responseIndex == 1 {
        return "Not a Jam!"
    }
    return "Jam!"
}