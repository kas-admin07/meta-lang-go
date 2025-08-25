package shortlinkapi

import (
	"net/http"
	"testing"
)

func TestShortenURL(t *testing.T) {
	httpGet := func(url string) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("shortened-url"))}, nil
	}
	http.DefaultClient.Get = httpGet
	result, err := ShortenURL("https://example.com")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != "shortened-url" {
		t.Errorf("Expected 'shortened-url', got %s", result)
	}
}
