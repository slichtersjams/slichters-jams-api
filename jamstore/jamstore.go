package jamstore

import (
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
	"strings"
)

type Jam struct {
	JamText string
	State bool
}

func StoreJam(ctx context.Context, jamText string, jamState bool) (error) {
	jam := new(Jam)
	jam.State = jamState
	jam.JamText = strings.ToLower(jamText)
	key := datastore.NewIncompleteKey(ctx, "Jam", nil)
	var err error
	key, err = datastore.Put(ctx, key, jam)

	return err
}
