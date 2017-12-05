package app

import "strings"

func storeUnknownJam(unknownJamStore IUnknownJamStore, jamText string) {
	lowerJamText := strings.ToLower(jamText)
	if unknownJamStore.GetJamKey(lowerJamText) == nil {
		unknownJamStore.StoreJam(lowerJamText)
	}
}

func clearUnknownJam(unknownJamStore IUnknownJamStore, jamText string) {
	lowerJamText := strings.ToLower(jamText)
	if key := unknownJamStore.GetJamKey(lowerJamText); key != nil {
		unknownJamStore.ClearJam(key)
	}
}
