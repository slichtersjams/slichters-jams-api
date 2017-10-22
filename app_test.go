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

func TestPostHandler__returns_bad_request_with_no_body(t *testing.T) {
	req, err := http.NewRequest("POST", "/jams", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(jamPostHandler)

	test_handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected %v, got %v", http.StatusBadRequest, status)
	}
}

func TestPostHandler__returns_bad_request_with_bad_json(t *testing.T) {
	reader := strings.NewReader("this is not json")
	req, err := http.NewRequest("POST", "/jams", reader)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(jamPostHandler)

	test_handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected %v, got %v", http.StatusBadRequest, status)
	}
}

func TestPostHandler__returns_bad_request_with_incorrect_json(t *testing.T) {
	reader := strings.NewReader(`{"Bar": "Foo"}`)
	req, err := http.NewRequest("POST", "/jams", reader)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(jamPostHandler)

	test_handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected %v, got %v", http.StatusBadRequest, status)
	}
}

func TestPostHandler__puts_jam_in_store_on_good_post_body_and_returns_200(t *testing.T) {
	inst, err := aetest.NewInstance(
		&aetest.Options{StronglyConsistentDatastore: true})

	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	reader := strings.NewReader(`{"JamText": "some jam text", "State": true}`)
	req, err := inst.NewRequest("POST", "/jams", reader)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(jamPostHandler)

	test_handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected 200, got %v", status)
	}
	jamText := "some jam text"
	query := datastore.NewQuery("Jam").Filter("JamText =", jamText)

	ctx := appengine.NewContext(req)
	var jams []Jam
	_, err = query.GetAll(ctx, &jams)
	if err != nil {
		t.Fatal(err)
	}

	if len(jams) == 0 {
		t.Fatal("No Jams found")
	}
	if query := jams[0].JamText; query != jamText {
		t.Errorf("Expected %v, got %v", jamText, query)
	}
	if state := jams[0].State; state != true {
		t.Errorf("Expected true, got %v", state)
	}
}