package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(apiHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Error("TestAPIGet failed: got: status %v, expected status 404", status)
	}
}

func TestAPIPost(t *testing.T) {
}
