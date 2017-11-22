package app

import "strings"

func storeUnknownJam(unknownJamStore IUnknownJamStore, jamText string) {
	lowerJamText := strings.ToLower(jamText)
	if !unknownJamStore.JamInStore(lowerJamText) {
		unknownJamStore.StoreJam(lowerJamText)
	}
}
