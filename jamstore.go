package app

import "strings"

type Jam struct {
	JamText string
	State bool
}

func StoreJam(data_store IDataStore, jamText string, jamState bool) (error) {
	jam := Jam{strings.ToLower(jamText), jamState}
	return data_store.Put(jam)
}

func GetJamState(data_store IDataStore, jamText string) bool {
	jam, _ := data_store.Get(jamText)

	return jam.State
}
