package app

type GifStore struct {

}

func (fake *GifStore)GetJamGif() string {
	return "https://media.giphy.com/media/l2QE6SbWP5RQKVVAc/giphy.gif"
}

func (fake *GifStore)GetNotJamGif() string {
	return "https://media.giphy.com/media/l2QEe1z9it3K4OeiY/giphy.gif"
}

func (fake *GifStore)GetVelourJamGif() string {
	return "https://media.giphy.com/media/xT0xeA9Njd9CzL8fyU/giphy.gif"
}
