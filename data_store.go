package app

import (
	"context"
	"google.golang.org/appengine/datastore"
)

type DataStore struct {
	Context context.Context
}

func (d *DataStore)Put(jam Jam) error {
	key := datastore.NewIncompleteKey(d.Context, "Jam", nil)
	_, err := datastore.Put(d.Context, key, &jam)

	return err
}

func (d *DataStore)Get(jamText string) (Jam, error) {
	query := datastore.NewQuery("Jam").Filter("JamText =", jamText)

	var jams []Jam
	_, err :=query.GetAll(d.Context, &jams)

	var jam Jam

	if len(jams) > 0 {
		jam = jams[0]
	} else if err == nil {
		err = datastore.ErrNoSuchEntity
	}

	return jam, err
}
