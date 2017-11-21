package app

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGifStore_GetJamGif(t *testing.T) {
	gifStore := new(GifStore)
	assert.Equal(t, "https://media.giphy.com/media/l2QE6SbWP5RQKVVAc/giphy.gif", gifStore.GetJamGif())
}

func TestGifStore_GetNotJamGif(t *testing.T) {
	gifStore := new(GifStore)
	assert.Equal(t, "https://media.giphy.com/media/l2QEe1z9it3K4OeiY/giphy.gif", gifStore.GetNotJamGif())
}

func TestGifStore_GetVelourJamGif(t *testing.T) {
	gifStore := new(GifStore)
	assert.Equal(t, "https://media.giphy.com/media/xT0xeA9Njd9CzL8fyU/giphy.gif", gifStore.GetVelourJamGif())
}
