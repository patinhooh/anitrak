package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anitrak",
	Short: "AniTrak is a CLI tool for tracking anime progress using AniList",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AniTrak CLI - Track your anime!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		// Handle error if needed
	}
}
