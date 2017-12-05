package app

import "google.golang.org/appengine/datastore"

type FakeUnknownJamStore struct {
	JamText string
	StoreCount int
	ClearCount int
}

func (fake *FakeUnknownJamStore)StoreJam(jamText string) {
	fake.JamText = jamText
	fake.StoreCount++
}

func (fake *FakeUnknownJamStore)GetJamKey(jamText string) *datastore.Key {
	if jamText == fake.JamText{
		return new(datastore.Key)
	}
	return nil
}

func (fake *FakeUnknownJamStore)ClearJam(key *datastore.Key) {
	fake.JamText = ""
	fake.ClearCount++
}
