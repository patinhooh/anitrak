package cmd

import (
	"fmt"

	"github.com/patinhooh/anitrak/internal/api"
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
		code, err := api.StartRedirectListener()
		if err != nil {
			fmt.Printf("Error capturing authorization code: %v\n", err)
			return
		}

		fmt.Printf("Authorization code received: %s\n", code)
		// Exchange the authorization code for an token
		err = api.ExchangeCodeForToken(code)
		if err != nil {
			fmt.Printf("Error exchanging code for token: %v\n", err)
			return
		}

		fmt.Println("Login successful! Saved Token")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
