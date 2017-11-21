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

func TestDataStore_PutUpdatesStateIfAlreadyInStore(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	assert.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("POST", "/", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(req)

	expectedJam := Jam{"some sweet jam text", true}
	ds := DataStore{ctx}

	key := datastore.NewIncompleteKey(ctx, "Jam", nil)
	_, err = datastore.Put(ctx, key, &expectedJam)
	assert.Nil(t, err)

	jam := Jam{"some sweet jam text", false}

	ds.Put(jam)

	query := datastore.NewQuery("Jam").Filter("JamText =", jam.JamText)

	var jams []Jam
	_, err = query.GetAll(ctx, &jams)
	assert.Nil(t, err)

	assert.Len(t, jams, 1)
}

func TestDataStore_Get(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	assert.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("POST", "/", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(req)

	expectedJam := Jam{"get some jams", true}
	ds := DataStore{ctx}

	key := datastore.NewIncompleteKey(ctx, "Jam", nil)
	_, err = datastore.Put(ctx, key, &expectedJam)
	assert.Nil(t, err)

	actualJam, _ := ds.Get("get some jams")

	assert.Equal(t, expectedJam, actualJam)
}

func TestDataStore_GetReturnsErrors(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	assert.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(req)

	expectedJam := Jam{"get some jams", true}
	ds := DataStore{ctx}

	key := datastore.NewIncompleteKey(ctx, "Jam", nil)
	_, err = datastore.Put(ctx, key, &expectedJam)
	assert.Nil(t, err)

	_, err = ds.Get("get some other jams")

	assert.NotNil(t, err)
}
