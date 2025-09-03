package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal("expected no error, got:", err)
	}
	if key != "my-secret-key" {
		t.Fatal("expected key to be my-secret-key, got:", key)
	}

	headers = http.Header{}
	_, err = GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error for missing Authorization header, got nil")
	}
}
