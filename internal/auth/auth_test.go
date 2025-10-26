package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	validApiKey := "test-api-key"
	testCases := []struct {
		headerKey string
		headerVal string
		isErr     bool
	}{
		{
			headerKey: "Authorization",
			headerVal: "ApiKey " + validApiKey,
			isErr:     false,
		},
		{
			headerKey: "Authorization",
			headerVal: "APIKEY " + validApiKey,
			isErr:     true,
		},
		{
			headerKey: "Authorization",
			headerVal: "ApiKey" + validApiKey,
			isErr:     true,
		},
		{
			headerKey: "Accept",
			headerVal: "application/json",
			isErr:     true,
		},
	}

	for _, c := range testCases {
		apiKey, err := auth.GetAPIKey(http.Header{c.headerKey: []string{c.headerVal}})
		if c.isErr {
			if err == nil {
				t.Errorf("GetAPIKey() error = %v, want error", err)
			}
		} else {
			if err != nil {
				t.Errorf("GetAPIKey() error = %v", err)
			}
			if apiKey != validApiKey {
				t.Errorf("GetAPIKey() = %v, want %v", apiKey, validApiKey)
			}
		}
	}
}
