package util

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/patinhooh/anitrak/internal/config"
	"github.com/patinhooh/anitrak/internal/models"
	"github.com/patinhooh/anitrak/internal/util"
	"github.com/stretchr/testify/assert"
)

func setupConfig() {
	config.TokenFilePath = filepath.Join(".anitrak", "config.json")
}

func cleanUp() {
	err := os.RemoveAll(filepath.Dir(config.TokenFilePath))
	if err != nil {
		log.Printf("Error cleaning up mock directory: %v", err)
	}
}

func TestMain(m *testing.M) {
	setupConfig()

	code := m.Run()

	cleanUp()
	os.Exit(code)
}

func TestSaveToken(t *testing.T) {
	token := &models.Token{
		Access:    "access_token",
		ExpiresAt: time.Now().Unix() + 3600,
		Type:      "bearer",
		Refresh:   "refresh_token",
	}

	err := util.SaveToken(token)
	assert.NoError(t, err)
}

func TestLoadToken(t *testing.T) {
	token, err := util.LoadToken()
	assert.NoError(t, err)
	assert.NotNil(t, token)
}

func TestIsTokenValid_JustExpired(t *testing.T) {
	// Prepare a token that expires at the exact moment of the test
	token := &models.Token{
		Access:    "access_token",
		ExpiresAt: time.Now().Unix(), // Token expires right now
		Type:      "bearer",
		Refresh:   "refresh_token",
	}

	// This might fail depending on how the time is handled by your logic,
	// but ideally the token should still be considered expired
	valid := util.IsTokenValid(token)
	assert.False(t, valid)
}

func TestIsTokenValid_Valid(t *testing.T) {
	// Prepare a token that expires in the future
	token := &models.Token{
		Access:    "access_token",
		ExpiresAt: time.Now().Unix() + 3600, // Token expires 1 hour from now
		Type:      "bearer",
		Refresh:   "refresh_token",
	}

	// The token should be valid
	valid := util.IsTokenValid(token)
	assert.True(t, valid)
}

func TestIsTokenValid_Invalid(t *testing.T) {
	// Prepare a token that expires in the future
	token := &models.Token{
		Access:    "access_token",
		ExpiresAt: time.Now().Unix() - 3600, // Token expires 1 hour from now
		Type:      "bearer",
		Refresh:   "refresh_token",
	}

	// The token should be valid
	valid := util.IsTokenValid(token)
	assert.False(t, valid)
}

func TestUpdateTokenData(t *testing.T) {
	tokenData := models.TokenData{
		Access:    "new_access_token",
		ExpiresIn: 7200,
		Type:      "bearer",
		Refresh:   "new_refresh_token",
	}
	token := &models.Token{}

	util.UpdateTokenData(tokenData, token)

	assert.Equal(t, token.Access, tokenData.Access)
	assert.Equal(t, token.ExpiresAt, time.Now().Unix()+tokenData.ExpiresIn)
	assert.Equal(t, token.Type, tokenData.Type)
	assert.Equal(t, token.Refresh, tokenData.Refresh)
}
