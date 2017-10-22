package app

import "google.golang.org/appengine/datastore"

type FakeDataStore struct {
	StoredJam Jam
	Error error
}

func (fake *FakeDataStore)Put(jam Jam) error {
	fake.StoredJam = jam
	return nil
}

func (fake *FakeDataStore)Get(jamText string) (Jam, error) {
	jam := Jam{"", false}
	err := fake.Error
	if err == nil {
		if jamText == fake.StoredJam.JamText {
			jam = fake.StoredJam
		} else {
			err = datastore.ErrNoSuchEntity
		}
	}
	return jam, err
}
