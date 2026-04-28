package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	// Status code test: expect 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Return value test: expect body "pong"
	if w.Body.String() != "pong" {
		t.Errorf("Expected body %q, got %q", "pong", w.Body.String())
	}

}
