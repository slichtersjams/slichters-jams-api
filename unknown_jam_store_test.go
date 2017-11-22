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

func TestFakeUnknownJamStore_JamInStore__returns_false_if_not_in_store(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	assert.Nil(t, err)
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(req)
	unknownJamStore := UnknownJamStore{ctx}

	assert.False(t, unknownJamStore.JamInStore("not in store"))
}