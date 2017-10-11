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
