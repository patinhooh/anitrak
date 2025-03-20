package config

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config struct to store all configuration values
type Config struct {
	GraphqlURL          string
	AniListClientId     string
	AniListClientSecret string
	AniListRedirectURL  string
	TokenURL            string
	AuthorizationURL    string
	TokenFilePath       string
}

func (cfg *Config) String() string {
	return fmt.Sprintf(
		"Config:\n  GraphqlURL: %s\n  AniListClientId: %s\n  AniListClientSecret: %s\n  AniListRedirectURL: %s\n  TokenURL: %s\n  AuthorizationURL: %s\n  TokenFilePath: %s",
		cfg.GraphqlURL, cfg.AniListClientId, cfg.AniListClientSecret, cfg.AniListRedirectURL, cfg.TokenURL, cfg.AuthorizationURL, cfg.TokenFilePath,
	)
}

func InitConfig() (*Config, error) {
	// TODO add path to this
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	config := &Config{}

	// Anilist Api
	config.GraphqlURL = os.Getenv("GRAPHQL_URL")
	if config.GraphqlURL == "" {
		return nil, fmt.Errorf("GRAPHQL_URL is not set in the .env file")
	}

	config.AniListClientId = os.Getenv("ANILIST_CLIENT_ID")
	if config.AniListClientId == "" {
		return nil, fmt.Errorf("ANILIST_CLIENT_ID is not set in the .env file")
	}

	config.AniListClientSecret = os.Getenv("ANILIST_CLIENT_SECRET")
	if config.AniListClientSecret == "" {
		return nil, fmt.Errorf("ANILIST_CLIENT_SECRET is not set in the .env file")
	}

	config.AniListRedirectURL = os.Getenv("ANILIST_REDIRECT_URL")
	if config.AniListRedirectURL == "" {
		return nil, fmt.Errorf("ANILIST_REDIRECT_URL is not set in the .env file")

	}

	// Anilist Auth
	config.TokenURL = "https://anilist.co/api/v2/oauth/token"

	config.AuthorizationURL = fmt.Sprintf(
		"https://anilist.co/api/v2/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s",
		config.AniListClientId,
		url.QueryEscape(config.AniListRedirectURL),
	)

	config.TokenFilePath = filepath.Join(os.Getenv("HOME"), ".anitrak", "client.json")

	return config, nil
}
