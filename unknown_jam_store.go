package app

import (
	"context"
	"google.golang.org/appengine/datastore"
)

type UnknownJam struct {
	JamText string
}

type UnknownJamStore struct {
	Context context.Context
}

func (store *UnknownJamStore)StoreJam(jamText string)  {
	key := datastore.NewIncompleteKey(store.Context, "UnknownJam", nil)
	datastore.Put(store.Context, key, &UnknownJam{JamText: jamText})
}

func (store *UnknownJamStore)JamInStore(jamText string) bool  {
	return false
}
