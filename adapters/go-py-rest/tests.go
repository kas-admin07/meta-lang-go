package gopyrest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallPythonService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("response"))
	}))
	defer server.Close()

	result, err := CallPythonService(server.URL, `{"key": "value"}`)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != "response" {
		t.Errorf("Expected 'response', got %s", result)
	}
}
