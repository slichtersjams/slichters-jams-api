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

	assert.Zero(t, fakeUnknownStore.StoreCount)
}

func TestStoreUnknownJam__ClearJam__removes_jam_from_unknown_jam_store(t *testing.T) {
	testText := "SoMe JaM tExT"

	fakeUnknownStore := new(FakeUnknownJamStore)
	fakeUnknownStore.JamText = strings.ToLower(testText)

	clearUnknownJam(fakeUnknownStore, testText)

	assert.Empty(t, fakeUnknownStore.JamText)
}

func TestStoreUnknownJam__ClearJam__removes_jam_from_unknown_jam_store_only_if_it_exists(t *testing.T) {
	testText := "SoMe JaM tExT"

	fakeUnknownStore := new(FakeUnknownJamStore)

	clearUnknownJam(fakeUnknownStore, testText)

	assert.Zero(t, fakeUnknownStore.ClearCount)
}
