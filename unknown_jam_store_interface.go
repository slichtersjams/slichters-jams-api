package app

type IUnknownJamStore interface {
	StoreJam(jamText string)
	JamInStore(jamText string) bool
}
