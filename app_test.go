package app

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	test_handler := http.HandlerFunc(handler)

	test_handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected 200, got %v", status)
	}

	expected := "Jam!"

	if actual := rr.Body.String(); actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
