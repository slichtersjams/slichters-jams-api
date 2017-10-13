package app

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/datastore"
)

func TestDataStore_Put(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	assert.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("POST", "/", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(req)

	ds := DataStore{ctx}

	jam := Jam{"some sweet jam text", true}

	ds.Put(jam)

	query := datastore.NewQuery("Jam").Filter("JamText =", jam.JamText)

	var jams []Jam
	_, err = query.GetAll(ctx, &jams)
	assert.Nil(t, err)

	assert.NotNil(t, jams)
}