package app

import "google.golang.org/appengine/datastore"

type IUnknownJamStore interface {
	StoreJam(jamText string)
	JamInStore(jamText string) *datastore.Key
}
