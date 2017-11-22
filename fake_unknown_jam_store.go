package app

type FakeUnknownJamStore struct {
	JamText string
	StoreCount int
}

func (fake *FakeUnknownJamStore)StoreJam(jamText string) {
	fake.JamText = jamText
	fake.StoreCount++
}

func (fake *FakeUnknownJamStore)JamInStore(jamText string) bool {
	return jamText == fake.JamText
}
