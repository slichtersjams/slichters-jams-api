package app

import "strings"

func storeUnknownJam(unknownJamStore IUnknownJamStore, jamText string) {
	lowerJamText := strings.ToLower(jamText)
	if unknownJamStore.JamInStore(lowerJamText) == nil {
		unknownJamStore.StoreJam(lowerJamText)
	}
}

func clearUnknownJam(unknownJamStore IUnknownJamStore, jamText string) {
	lowerJamText := strings.ToLower(jamText)
	unknownJamStore.ClearJam(lowerJamText)
}
