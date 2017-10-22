package app

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

type FakeDataStore struct {
	StoredJam Jam
}

func (fake *FakeDataStore)Put(jam Jam) error {
	fake.StoredJam = jam
	return nil
}

func (fake *FakeDataStore)Get(jamText string) (Jam, error) {
	jam := Jam{"", false}
	if jamText == fake.StoredJam.JamText {
		jam = fake.StoredJam
	}
	return jam, nil
}

func TestStoreJam__correctly_stores_jam_in_datastore(t *testing.T) {
	fakeDataStore := new(FakeDataStore)

	expectedJam := Jam{"foo", true}
	err := StoreJam(fakeDataStore, expectedJam.JamText, expectedJam.State)
	assert.Nil(t, err)

	assert.Equal(t, expectedJam, fakeDataStore.StoredJam)
}

func TestStoreJam__correctly_stores_jam_in_datastore_with_lower_case_text(t *testing.T) {
	fakeDataStore := new(FakeDataStore)

	expectedJamText := "foo"
	err := StoreJam(fakeDataStore, strings.ToUpper(expectedJamText), true)
	assert.Nil(t, err)

	assert.Equal(t, expectedJamText, fakeDataStore.StoredJam.JamText)
}

func TestGetJam__gets_jam_from_datastore(t *testing.T) {
	storedJam := Jam{"meat loaves", true}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam
	jamState := GetJamState(fakeDataStore, storedJam.JamText)
	assert.Equal(t, storedJam.State, jamState)
}

func TestGetJam__gets_jam_from_datastore_when_text_is_not_lower_case(t *testing.T) {
	storedJam := Jam{"meat loaves", true}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam
	jamState := GetJamState(fakeDataStore, "MeAt LoAvEs")
	assert.Equal(t, storedJam.State, jamState)
}
