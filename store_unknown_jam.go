package app

func storeUnknownJam(unknownJamStore IUnknownJamStore, jamText string) {
	unknownJamStore.StoreJam(jamText)
}
