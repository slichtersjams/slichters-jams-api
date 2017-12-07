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

func (store *UnknownJamStore)GetJamKey(jamText string) *datastore.Key  {
	query := datastore.NewQuery("UnknownJam").Filter("JamText =", jamText)

	var unknownJams []UnknownJam
	keys, _ := query.GetAll(store.Context, &unknownJams)
	if len(keys) > 0 {
		return keys[0]
	}
	return nil
}

func (store *UnknownJamStore)ClearJam(key *datastore.Key) {
	datastore.Delete(store.Context, key)
}

func (store *UnknownJamStore)GetAllJams() []string {
	query := datastore.NewQuery("UnknownJam")
	var unknownJams []UnknownJam
	query.GetAll(store.Context, &unknownJams)

	jamList := make([]string, len(unknownJams))

	for index, jam := range unknownJams {
		jamList[index] = jam.JamText
	}

	return jamList
}
