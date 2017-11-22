package app

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type FakeUnknownJamStore struct {
	JamText string
}

func (fake *FakeUnknownJamStore)StoreJam(jamText string) {
	fake.JamText = jamText
}

func TestStoreUnknownJam__puts_jam_in_unknown_store(t *testing.T)  {
	fakeUnknownStore := new(FakeUnknownJamStore)

	testText := "some jam text"
	storeUnknownJam(fakeUnknownStore, testText)

	assert.Equal(t, testText, fakeUnknownStore.JamText)
}
