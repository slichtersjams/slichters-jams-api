package app

type IGifStore interface {
	GetJamGif() string
	GetNotJamGif() string
	GetVelourJamGif() string
}
