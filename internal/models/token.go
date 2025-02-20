package models

type TokenData struct {
	Access    string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
	Type      string `json:"token_type"`
	Refresh   string `json:"refresh_token"`
}

type Token struct {
	Access    string
	ExpiresAt int64
	Type      string
	Refresh   string
}
