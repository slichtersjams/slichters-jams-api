package app

import "google.golang.org/appengine/datastore"

type FakeUnknownJamStore struct {
	JamText string
	StoreCount int
}

func (fake *FakeUnknownJamStore)StoreJam(jamText string) {
	fake.JamText = jamText
	fake.StoreCount++
}

func (fake *FakeUnknownJamStore)JamInStore(jamText string) *datastore.Key {
	if jamText == fake.JamText{
		return new(datastore.Key)
	}
	return nil
}
