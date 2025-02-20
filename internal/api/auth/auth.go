package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patinhooh/anitrak/internal/api/auth"
	"github.com/patinhooh/anitrak/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStartRedirectListener(t *testing.T) {
	// Create a mock HTTP request with a query parameter "code"
	req := httptest.NewRequest("GET", "/callback?code=test_code", nil)
	rr := httptest.NewRecorder()

	// Create a test handler for the callback URL
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		assert.Equal(t, "test_code", code)
		w.Write([]byte("Authorization successful!"))
	})

	// Simulate the request
	http.DefaultServeMux.ServeHTTP(rr, req)

	// Assert the status code and the response body
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Authorization successful!", rr.Body.String())
}

func TestExchangeCodeForToken(t *testing.T) {
	// Mock HTTP server to simulate AniList token endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assert POST request and URL
		assert.Equal(t, "/token", r.URL.Path)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))

		// Mock successful response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token": "access_token_value", "refresh_token": "refresh_token_value"}`))
	}))
	defer ts.Close()

	// Mock AniList token URL in your config to point to the test server
	auth.Config.TokenURL = ts.URL + "/token"

	// Call the function with a sample authorization code
	err := auth.ExchangeCodeForToken("test_code")
	assert.NoError(t, err)
}

func TestRefreshToken(t *testing.T) {
	// Mock HTTP server to simulate AniList refresh token endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assert POST request and URL
		assert.Equal(t, "/token", r.URL.Path)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))

		// Mock successful refresh response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token": "new_access_token", "refresh_token": "new_refresh_token"}`))
	}))
	defer ts.Close()

	// Mock AniList token URL in your config to point to the test server
	auth.Config.TokenURL = ts.URL + "/token"

	// Example token to refresh
	token := &models.Token{Refresh: "existing_refresh_token"}

	// Call the function with the mock token
	err := auth.RefreshToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "new_access_token", token.Access)
	assert.Equal(t, "new_refresh_token", token.Refresh)
}
