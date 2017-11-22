package app

import "strings"

func storeUnknownJam(unknownJamStore IUnknownJamStore, jamText string) {
	unknownJamStore.StoreJam(strings.ToLower(jamText))
}
