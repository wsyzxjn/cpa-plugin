package main

import (
	"encoding/base64"
	"net/http"
	"testing"
)

func TestDecodeHostHTTPResponse(t *testing.T) {
	body := []byte(`{"code":0}`)
	encodedBody := base64.StdEncoding.EncodeToString(body)

	tests := []struct {
		name string
		raw  string
	}{
		{
			name: "current host field",
			raw:  `{"StatusCode":200,"Headers":{"Content-Type":["application/json"]},"Body":"` + encodedBody + `"}`,
		},
		{
			name: "snake case field",
			raw:  `{"status_code":200,"headers":{"Content-Type":["application/json"]},"body":"` + encodedBody + `"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := decodeHostHTTPResponse([]byte(test.raw))
			if err != nil {
				t.Fatalf("decode response: %v", err)
			}
			if response.StatusCode != http.StatusOK {
				t.Fatalf("status code = %d, want %d", response.StatusCode, http.StatusOK)
			}
			if response.Headers.Get("Content-Type") != "application/json" {
				t.Fatalf("content type = %q", response.Headers.Get("Content-Type"))
			}
			if string(response.Body) != string(body) {
				t.Fatalf("body = %q, want %q", response.Body, body)
			}
		})
	}
}
