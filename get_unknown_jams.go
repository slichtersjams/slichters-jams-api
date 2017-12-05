package app

import (
	"net/http"
	"encoding/json"
)

type UnknownJamJson struct {
	UnknownJams []string
}

func GetUnknownJams(unknownJamStore IUnknownJamStore, w http.ResponseWriter) {
	unknown_jam_list := unknownJamStore.GetAllJams()
	response := UnknownJamJson{unknown_jam_list}
	js, _ := json.Marshal(response)
	w.Write(js)
}
