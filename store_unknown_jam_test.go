package app

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestStoreUnknownJam__puts_jam_in_unknown_store(t *testing.T)  {
	fakeUnknownStore := new(FakeUnknownJamStore)

	testText := "some jam text"
	storeUnknownJam(fakeUnknownStore, testText)

	assert.Equal(t, testText, fakeUnknownStore.JamText)
}

func TestStoreUnknownJam__puts_jam_in_unknown_store_in_lower_case(t *testing.T) {
	fakeUnknownStore := new(FakeUnknownJamStore)

	testText := "SoMe JaM tExT"
	storeUnknownJam(fakeUnknownStore, testText)

	assert.Equal(t, strings.ToLower(testText), fakeUnknownStore.JamText)
}

func TestStoreUnknownJam__puts_jam_in_unknown_store_if_not_already_there(t *testing.T) {
	testText := "SoMe JaM tExT"

	fakeUnknownStore := new(FakeUnknownJamStore)
	fakeUnknownStore.JamText = strings.ToLower(testText)

	storeUnknownJam(fakeUnknownStore, testText)

	assert.Equal(t, 0, fakeUnknownStore.StoreCount)
}