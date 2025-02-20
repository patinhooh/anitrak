package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/patinhooh/anitrak/internal/config"
	"github.com/patinhooh/anitrak/internal/models"
	"github.com/patinhooh/anitrak/internal/util"
)

// StartRedirectListener starts a small HTTP server to capture the OAuth code
func StartRedirectListener() (string, error) {
	var code string
	server := &http.Server{Addr: ":8080"}

	// Handler to capture the code from the query parameter
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Get the code from the query parameters
		code = r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Authorization code not found", http.StatusBadRequest)
			return
		}

		// Respond to the user
		fmt.Fprintf(w, "Authorization successful! You can return to the CLI now.")

		// Shut down the server once the code is captured
		go func() {
			_ = server.Close()
		}()
	})

	// Start the server
	fmt.Println("Listening on http://localhost:8080 for the redirect...")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return "", fmt.Errorf("failed to start HTTP server: %v", err)
	}

	return code, nil
}

func ExchangeCodeForToken(code string) error {
	data := url.Values{}
	data.Set("client_id", config.AniListClientId)
	data.Set("client_secret", config.AniListClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", config.AniListRedirectURI)
	data.Set("grant_type", "authorization_code")

	// Make the request to AniList's token endpoint
	resp, err := http.PostForm(config.TokenURL, data)
	if err != nil {
		return fmt.Errorf("failed to make token request: %v", err)
	}
	defer resp.Body.Close()

	// Debug: Check the response status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	// Debug: Read and log the response body
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Raw response body: %s\n", string(body))

	// Parse the response
	var tokenData models.TokenData
	var token models.Token
	if err := json.Unmarshal(body, &tokenData); err != nil {
		return fmt.Errorf("failed to parse token response: %v", err)
	}
	util.UpdateTokenData(tokenData, &token)

	// Save Token
	err = util.SaveToken(&token)
	if err != nil {
		return fmt.Errorf("failed to save token: %v", err)
	}

	return nil
}

func RefreshToken(token *models.Token) error {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", config.AniListClientId)
	data.Set("client_secret", config.AniListClientSecret)
	data.Set("refresh_token", token.Refresh)

	// Make the request
	resp, err := http.PostForm(config.TokenURL, data)
	if err != nil {
		return fmt.Errorf("failed to make refresh token request: %v", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to refresh token: %s", string(body))
	}

	// Parse the response
	var newTokenData models.TokenData
	if err := json.NewDecoder(resp.Body).Decode(&newTokenData); err != nil {
		return fmt.Errorf("failed to parse refresh token response: %v", err)
	}

	util.UpdateTokenData(newTokenData, token)

	// Save the updated token
	if err := util.SaveToken(token); err != nil {
		return fmt.Errorf("failed to save updated token: %v", err)
	}

	fmt.Println("Token refreshed successfully!")
	return nil
}
