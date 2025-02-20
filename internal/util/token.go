package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/patinhooh/anitrak/internal/config"
	"github.com/patinhooh/anitrak/internal/models"
)

func SaveToken(token *models.Token) error {
	dir := filepath.Dir(config.TokenFilePath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	file, err := os.OpenFile(config.TokenFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to save token: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(token); err != nil {
		return fmt.Errorf("failed to write token data: %v", err)
	}
	return nil
}

func LoadToken() (*models.Token, error) {
	// Open the token file
	file, err := os.Open(config.TokenFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load token: %v", err)
	}
	defer file.Close()

	// Decode the JSON into Token
	var token models.Token
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&token); err != nil {
		return nil, fmt.Errorf("failed to parse token data: %v", err)
	}

	return &token, nil
}

func IsTokenValid(token *models.Token) bool {
	return token.ExpiresAt > time.Now().Unix()
}

func UpdateTokenData(tokenData models.TokenData, token *models.Token) {
	token.Access = tokenData.Access
	token.ExpiresAt = time.Now().Unix() + tokenData.ExpiresIn
	token.Type = tokenData.Type
	token.Refresh = tokenData.Refresh
}
