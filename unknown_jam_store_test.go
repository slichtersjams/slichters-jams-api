package app

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/datastore"
)

func TestUnknownJamStore_StoreJam(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	assert.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(req)

	unknownJamStore := UnknownJamStore{ctx}

	testText := "some jam text"
	unknownJamStore.StoreJam(testText)

	query := datastore.NewQuery("UnknownJam").Filter("JamText =", testText)

	var unknownJams []UnknownJam
	_, err = query.GetAll(ctx, &unknownJams)
	assert.Nil(t, err)

	assert.NotNil(t, unknownJams)
}

func TestUnknownJamStore_JamInStore__returns_nil_if_not_in_store(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	assert.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(req)
	unknownJamStore := UnknownJamStore{ctx}

	assert.Nil(t, unknownJamStore.JamInStore("not in store"))
}

func TestUnknownJamStore_JamInStore__returns_key_if_in_store(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	assert.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(req)

	testText := "jam in store"
	key := datastore.NewIncompleteKey(ctx, "UnknownJam", nil)
	_, err = datastore.Put(ctx, key, &UnknownJam{JamText: testText})
	assert.Nil(t, err)

	unknownJamStore := UnknownJamStore{ctx}

	assert.NotNil(t, unknownJamStore.JamInStore(testText))
}
