package jamstore

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
)

func TestStoreJam__correctly_stores_jam_in_datastore(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := appengine.NewContext(req)

	jamQuery := "foo"
	jamState := true
	if err := StoreJam(ctx, jamQuery, jamState); err != nil {
		t.Fatal(err)
	}

	query := datastore.NewQuery("Jam").Filter("JamText =", jamQuery)

	var jams []Jam
	_, err = query.GetAll(ctx, &jams)
	if err != nil {
		t.Fatal(err)
	}

	if len(jams) == 0 {
		t.Fatal("No Jams found")
	}
	if query := jams[0].JamText; query != jamQuery {
		t.Errorf("Expected %v, got %v", jamQuery, query)
	}
	if state := jams[0].State; state != jamState {
		t.Errorf("Expected %v, got %v", jamState, state)
	}
}
