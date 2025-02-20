package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// var AniListAPI string
var AniListClientId string
var AniListClientSecret string
var AniListRedirectURI string

var TokenURL string
var AuthorizationURL string
var TokenFilePath string

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// // Anilist Api
	// AniListAPI = os.Getenv("ANILIST_API")
	// if AniListAPI == "" {
	// 	log.Fatalf("ANILIST_API is not set in the .env file")
	// }

	AniListClientId = os.Getenv("ANILIST_CLIENT_ID")
	if AniListClientId == "" {
		log.Fatalf("ANILIST_CLIENT_ID is not set in the .env file")
	}

	AniListClientSecret = os.Getenv("ANILIST_CLIENT_SECRET")
	if AniListClientSecret == "" {
		log.Fatalf("ANILIST_CLIENT_SECRET is not set in the .env file")
	}

	AniListRedirectURI = os.Getenv("ANILIST_REDIRECT_URI")
	if AniListRedirectURI == "" {
		log.Fatalf("ANILIST_REDIRECT_URI is not set in the .env file")
	}

	// Anilist Auth
	TokenURL = "https://anilist.co/api/v2/oauth/token"

	AuthorizationURL = fmt.Sprintf(
		"https://anilist.co/api/v2/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s",
		AniListClientId,
		url.QueryEscape(AniListRedirectURI),
	)

	TokenFilePath = filepath.Join(os.Getenv("HOME"), ".anitrak", "config.json")
}
