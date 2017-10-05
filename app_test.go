package app

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestHandler__returns_correct_responses(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(getHandler)

	test_handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected 200, got %v", status)
	}

	expected := "Jam!"

	if actual := rr.Body.String(); actual != expected {
		expected := "Not a Jam!"
		if actual := rr.Body.String(); actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}

	expected_header := "*"
	if actual_header := rr.Header().Get("Access-Control-Allow-Origin"); actual_header != expected_header {
		t.Errorf("Expected %v, got %v", expected_header, actual_header)
	}
}

func TestGetResponse__returns_jam_for_zero(t *testing.T) {
	expected := "Jam!"
	if resp := getResponse(0); resp != expected {
		t.Errorf("Expected %v, got %v", expected, resp)
	}
}

func TestGetResponse__returns_not_a_jam_for_1(t *testing.T) {
	expected := "Not a Jam!"
	if resp := getResponse(1); resp != expected {
		t.Errorf("Expected %v, got %v", expected, resp)
	}
}

//func TestPostHandler__returns_200(t *testing.T) {
//	req, err := http.NewRequest("POST", "/jams", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	test_handler := http.HandlerFunc(jamPostHandler)
//
//	test_handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("Expected 200, got %v", status)
//	}
//}

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
