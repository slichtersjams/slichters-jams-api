package app

import (
	"testing"
	"net/http/httptest"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"fmt"
)

func TestGetUnknownJams__has_empty_unknown_jam_list_if_none_in_store(t *testing.T) {
	rr := httptest.NewRecorder()
	fakeUnknownJamStore := new(FakeUnknownJamStore)

	GetUnknownJams(fakeUnknownJamStore, rr)
	assert.Equal(t, http.StatusOK, rr.Code)

	fmt.Println(rr.Body)
	decoder := json.NewDecoder(rr.Body)
	var response UnknownJamJson
	err := decoder.Decode(&response)
	assert.Nil(t, err)
	assert.Empty(t, response.UnknownJams)
}

func TestGetUnknownJams__has_jam_the_correct_jam_list(t *testing.T) {
	rr := httptest.NewRecorder()
	fakeUnknownJamStore := new(FakeUnknownJamStore)
	jam_list := [2]string{"some jam", "some other jam"}
	fakeUnknownJamStore.AllJams = jam_list[:]

	GetUnknownJams(fakeUnknownJamStore, rr)

	assert.Equal(t, http.StatusOK, rr.Code)

	fmt.Println(rr.Body)
	decoder := json.NewDecoder(rr.Body)
	var response UnknownJamJson
	err := decoder.Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, jam_list[:], response.UnknownJams)
}
