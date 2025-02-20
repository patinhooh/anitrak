package cmd

import (
	"fmt"

	"github.com/patinhooh/anitrak/internal/api/auth"
	"github.com/patinhooh/anitrak/internal/config"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to your AniList account",
	Run: func(cmd *cobra.Command, args []string) {
		// Redirect user to authorization URL
		fmt.Println("Please visit the following URL to authorize your application:")
		fmt.Println(config.AuthorizationURL)

		// Start Server that capture the callback
		code, err := auth.StartRedirectListener()
		if err != nil {
			fmt.Printf("Error capturing authorization code: %v\n", err)
			return
		}

		fmt.Printf("Authorization code received: %s\n", code)
		// Exchange the authorization code for an access token
		err = auth.ExchangeCodeForToken(code)
		if err != nil {
			fmt.Printf("Error exchanging code for token: %v\n", err)
			return
		}

		// Store the access token
		fmt.Println("Login successful! Saved Token")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
