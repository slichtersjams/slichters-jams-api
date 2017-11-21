package app

type FakeGifStore struct {
	JamGif string
	NotJamGif string
	VelourJamGif string
}

func (fake *FakeGifStore)GetJamGif() string {
	return fake.JamGif
}

func (fake *FakeGifStore)GetNotJamGif() string {
	return fake.NotJamGif
}

func (fake *FakeGifStore)GetVelourJamGif() string {
	return fake.VelourJamGif
}
