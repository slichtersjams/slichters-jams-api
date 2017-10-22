package app

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"github.com/stretchr/testify/assert"
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

func TestGetJamResponse__returns_correct_response_if_it_is_a_jam(t *testing.T) {
	storedJam := Jam{"meat loaves", true}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, "meat loaves", rr)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Jam!", rr.Body.String())
}

func TestGetJamResponse__returns_correct_response_if_it_is_not_a_jam(t *testing.T) {
	storedJam := Jam{"meat loaves", false}
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.StoredJam = storedJam

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, "meat loaves", rr)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Not a Jam!", rr.Body.String())
}

func TestGetJamResponse__returns_bad_request_if_query_not_in_data_store(t *testing.T) {
	fakeDataStore := new(FakeDataStore)

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, "meat loaves", rr)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetJamResponse__returns_internal_error_if_datastore_has_errors(t *testing.T) {
	fakeDataStore := new(FakeDataStore)
	fakeDataStore.Error = datastore.ErrInvalidKey

	rr := httptest.NewRecorder()

	getJamResponse(fakeDataStore, "meat loaves", rr)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Internal Server Error : datastore: invalid key\n", rr.Body.String())
}

func TestPostHandler__returns_bad_request_with_no_body(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	assert.Nil(t, err)

	defer inst.Close()

	req, err := inst.NewRequest("POST", "/jams", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(jamPostHandler)

	test_handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostHandler__returns_bad_request_with_bad_json(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	assert.Nil(t, err)

	defer inst.Close()

	reader := strings.NewReader("this is not json")
	req, err := inst.NewRequest("POST", "/jams", reader)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(jamPostHandler)

	test_handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostHandler__returns_bad_request_with_incorrect_json(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})
	assert.Nil(t, err)

	defer inst.Close()

	reader := strings.NewReader(`{"Bar": "Foo"}`)
	req, err := inst.NewRequest("POST", "/jams", reader)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(jamPostHandler)

	test_handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostHandler__puts_jam_in_store_on_good_post_body_and_returns_200(t *testing.T) {
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

	jamText := "some jam text"
	query := datastore.NewQuery("Jam").Filter("JamText =", jamText)

	ctx := appengine.NewContext(req)
	var jams []Jam
	_, err = query.GetAll(ctx, &jams)
	assert.Nil(t, err)

	assert.NotEmpty(t, jams)
	assert.Equal(t, jamText, jams[0].JamText)
	assert.Equal(t, true, jams[0].State)
}