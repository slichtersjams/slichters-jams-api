package app

import (
	"google.golang.org/appengine/datastore"
	"strings"
)

type Jam struct {
	JamText string
	State bool
}

func StoreJam(ctx IGetContext, jamText string, jamState bool) (error) {
	jam := new(Jam)
	jam.State = jamState
	jam.JamText = strings.ToLower(jamText)
	context := ctx.getContext()
	key := datastore.NewIncompleteKey(context, "Jam", nil)
	var err error
	key, err = datastore.Put(context, key, jam)

	return err
}
