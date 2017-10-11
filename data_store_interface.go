package app

type IDataStore interface {
	Put(jam Jam) error
}
