package app

import "google.golang.org/appengine/datastore"

type IUnknownJamStore interface {
	StoreJam(jamText string)
	GetJamKey(jamText string) *datastore.Key
	ClearJam(key *datastore.Key)
	GetAllJams() []string
}
