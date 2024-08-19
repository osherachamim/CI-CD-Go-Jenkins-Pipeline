package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/World", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	HelloServer(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", rec.Code)
	}

	expected := "Hello, World!"
	if rec.Body.String() != expected {
		t.Fatalf("Expected %q; got %q", expected, rec.Body.String())
	}
}
