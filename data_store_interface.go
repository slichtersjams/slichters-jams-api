package app

type IDataStore interface {
	Put(jam Jam) error
	Get(jamText string) (Jam, error)
}
