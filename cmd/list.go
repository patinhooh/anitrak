package cmd

import (
	"github.com/patinhooh/anitrak/internal/api"
	"github.com/spf13/cobra"
)

var animeListCmd = &cobra.Command{
	Use:   "anime list",
	Short: "Log in to your AniList account",
	Run: func(cmd *cobra.Command, args []string) {

		// Print out the response
		api.PrintAnimeList()
	},
}

func init() {
	rootCmd.AddCommand(animeListCmd)
}
