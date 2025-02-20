package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anitrak",
	Short: "AniTrak is a CLI client for AniList API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("                _ _______        _")
		fmt.Println("    /\\         (_)__   __|      | |")
		fmt.Println("   /  \\   _ __  _   | |_ __ __ _| | __")
		fmt.Println("  / /\\ \\ | '_ \\| |  | | '__/ _` | |/ /")
		fmt.Println(" / ____ \\| | | | |  | | | | (_| |   < ")
		fmt.Println("/_/    \\_\\_| |_|_|  |_|_|  \\__,_|_|\\_\\")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
