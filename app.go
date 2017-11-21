package app

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
    "encoding/json"
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)

type ResponseJson struct {
	JamState bool
	JamText string
	JamGif string
}

var GetRandomJam = getRandomResponse

func init() {
    http.HandleFunc("/", getHandler)
    http.HandleFunc("/jams", jamPostHandler)
    rand.Seed(time.Now().UTC().UnixNano())
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    jamText := r.URL.Query().Get("jamText")
    if jamText != "" {
        ctx := appengine.NewContext(r)
        dataStore := DataStore{ctx}

        getJamResponse(&dataStore, jamText, w)
    } else {
        fmt.Fprint(w, GetRandomJam())
    }
}

func getJamResponse(dataStore IDataStore, jamText string, w http.ResponseWriter) {
    jamState, err := GetJamState(dataStore, jamText)
    if err != nil {
        if err == datastore.ErrNoSuchEntity {
            http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        } else {
            http.Error(w, http.StatusText(http.StatusInternalServerError) + " : " + err.Error(),
                http.StatusInternalServerError)
        }
    } else {
        response := ResponseJson{JamGif: "https://media0.giphy.com/media/3otPovEi2MtN9pEJuo/giphy.gif"}
        response.JamState = jamState
        if jamState {
            response.JamText = "Jam"
        } else {
            response.JamText = "NotJam"
        }
		js, _ := json.Marshal(response)
		w.Write(js)
    }
    w.Header().Set("Content-Type", "application/json")
}

func jamPostHandler(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)

    dataStore := DataStore{ctx}

    postJam(r, w, &dataStore)
}

func postJam(r *http.Request, w http.ResponseWriter, dataStore IDataStore) {
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
    if err = StoreJam(dataStore, jam.JamText, jam.State); err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    }
}

func getRandomResponse() string  {
    if rand.Intn(2) == 1 {
        return "Not a Jam!"
    }
    return "Jam!"
}