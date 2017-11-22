package app

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func fakeRandomJamGenerator() string {
	return "Random Jam"
}

func TestHandler__returns_random_response_if_no_query(t *testing.T) {
	var oldRandJamFunc = GetRandomJam
	GetRandomJam = fakeRandomJamGenerator
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(getHandler)

	test_handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Random Jam", rr.Body.String())
	assert.Equal(t, "*", rr.Header().Get("Access-Control-Allow-Origin"))
	GetRandomJam = oldRandJamFunc
}

func TestGetJamResponse__puts_json_content_type_header(t *testing.T) {
	storedJam := Jam{"meat loaves", true}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam

	fakeGifStore := new(FakeGifStore)
	fakeUnknownJamStore := new(FakeUnknownJamStore)

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore, "meat loaves", rr)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestGetJamResponse__returns_correct_response_if_it_is_a_jam(t *testing.T) {
	storedJam := Jam{"meat loaves", true}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam

	fakeGifStore := new(FakeGifStore)
	fakeUnknownJamStore := new(FakeUnknownJamStore)

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore,"meat loaves", rr)
	assert.Equal(t, http.StatusOK, rr.Code)

	decoder := json.NewDecoder(rr.Body)
	var response ResponseJson
	err := decoder.Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, "Jam", response.JamText)
	assert.Equal(t, true, response.JamState)
}

func TestGetJamResponse__returns_correct_response_if_it_is_not_a_jam(t *testing.T) {
	storedJam := Jam{"meat loaves", false}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam

	fakeGifStore := new(FakeGifStore)
	fakeUnknownJamStore := new(FakeUnknownJamStore)

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore,"meat loaves", rr)

	decoder := json.NewDecoder(rr.Body)
	var response ResponseJson
	err := decoder.Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, "NotJam", response.JamText)
	assert.Equal(t, false, response.JamState)
}

func TestGetJamResponse__returns_not_jam_gif_link(t *testing.T) {
	storedJam := Jam{"meat loaves", false}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam

	fakeGifStore := new(FakeGifStore)
	fakeGifStore.NotJamGif = "not a jam gif"

	fakeUnknownJamStore := new(FakeUnknownJamStore)

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore,"meat loaves", rr)

	decoder := json.NewDecoder(rr.Body)
	var response ResponseJson
	err := decoder.Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, fakeGifStore.NotJamGif, response.JamGif)
}

func TestGetJamResponse__returns_velour_gif_link_if_given_velour_text(t *testing.T) {
	storedJam := Jam{"velour tracksuit", true}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam

	fakeGifStore := new(FakeGifStore)
	fakeGifStore.VelourJamGif = "velour gif"

	fakeUnknownJamStore := new(FakeUnknownJamStore)

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore,"velour tracksuit", rr)

	decoder := json.NewDecoder(rr.Body)
	var response ResponseJson
	err := decoder.Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, fakeGifStore.VelourJamGif, response.JamGif)
}

func TestGetJamResponse__returns_jam_gif_link(t *testing.T) {
	storedJam := Jam{"meat loaves", true}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam

	fakeGifStore := new(FakeGifStore)
	fakeGifStore.JamGif = "some jam gif"

	fakeUnknownJamStore := new(FakeUnknownJamStore)

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore,"meat loaves", rr)

	decoder := json.NewDecoder(rr.Body)
	var response ResponseJson
	err := decoder.Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, fakeGifStore.JamGif, response.JamGif)
}

func TestGetJamResponse__returns_bad_request_if_query_not_in_data_store(t *testing.T) {
	fakeDataStore := new(FakeDataStore)

	fakeGifStore := new(FakeGifStore)

	rr := httptest.NewRecorder()
	fakeUnknownJamStore := new(FakeUnknownJamStore)

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore,"meat loaves", rr)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetJamResponse__puts_jam_in_unknown_store_if_query_not_in_data_store(t *testing.T) {
	fakeDataStore := new(FakeDataStore)

	fakeGifStore := new(FakeGifStore)

	rr := httptest.NewRecorder()
	fakeUnknownJamStore := new(FakeUnknownJamStore)

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore,"Meat Loaves", rr)
	assert.Equal(t,"meat loaves", fakeUnknownJamStore.JamText)
}

func TestGetJamResponse__returns_internal_error_if_datastore_has_errors(t *testing.T) {
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.Error = datastore.ErrInvalidKey

	fakeGifStore := new(FakeGifStore)
	fakeUnknownJamStore := new(FakeUnknownJamStore)

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, fakeGifStore, fakeUnknownJamStore, "meat loaves", rr)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Internal Server Error : datastore: invalid key\n", rr.Body.String())
}

func TestPostJam__returns_bad_request_with_no_body(t *testing.T) {
	req, err := http.NewRequest("POST", "/jams", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	fakeDataStore := new(FakeDataStore)

	postJam(req, rr, fakeDataStore)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostJam__returns_bad_request_with_bad_json(t *testing.T) {
	reader := strings.NewReader("this is not json")
	req, err := http.NewRequest("POST", "/jams", reader)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	fakeDataStore := new(FakeDataStore)

	postJam(req, rr, fakeDataStore)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostJam__returns_bad_request_with_incorrect_json(t *testing.T) {
	reader := strings.NewReader(`{"Bar": "Foo"}`)
	req, err := http.NewRequest("POST", "/jams", reader)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	fakeDataStore := new(FakeDataStore)

	postJam(req, rr, fakeDataStore)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostJam__returns_internal_server_failure_when_datastore_has_errors(t *testing.T) {
	reader := strings.NewReader(`{"JamText": "some jam text", "State": true}`)
	req, err := http.NewRequest("POST", "/jams", reader)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	fakeDataStore := new(FakeDataStore)
	fakeDataStore.Error = datastore.ErrInvalidKey

	postJam(req, rr, fakeDataStore)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestPostJam__puts_jam_in_store(t *testing.T) {
	reader := strings.NewReader(`{"JamText": "some jam text", "State": true}`)
	req, err := http.NewRequest("POST", "/jams", reader)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	fakeDataStore := new(FakeDataStore)

	postJam(req, rr, fakeDataStore)

	expectedJam := Jam{"some jam text", true}

	assert.Equal(t, expectedJam, fakeDataStore.StoredJam)
}

func TestPostHandler__returns_200_with_good_json(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	assert.Nil(t, err)
	defer inst.Close()

	reader := strings.NewReader(`{"JamText": "some jam text", "State": true}`)
	req, err := inst.NewRequest("POST", "/jams", reader)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(jamPostHandler)

	test_handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}