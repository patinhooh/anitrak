package cmd

import (
	"fmt"
	"os"

	"github.com/patinhooh/anitrak/internal/config"
	"github.com/spf13/cobra"
)

// Used for flags.
var cfg *config.Config
var rootCmd = &cobra.Command{
	Use:   "anitrak",
	Short: "AniTrak is a CLI client for AniList API",
	Long:  `AniTrak is a CLI client for AniList API`,
	// TODO make the long version
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, err = config.InitConfig()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("                _ \033[35m_______        _\033[0m")
		fmt.Println("    /\\         (_)\033[35m__   __|      | |\033[0m")
		fmt.Println("   /  \\   _ __  _   \033[35m| |_ __ __ _| | __\033[0m")
		fmt.Println("  / /\\ \\ | '_ \\| |  \033[35m| | '__/ _` | |/ /\033[0m")
		fmt.Println(" / ____ \\| | | | |  \033[35m| | | | (_| |   < \033[0m")
		fmt.Println("/_/    \\_\\_| |_|_|  \033[35m|_|_|  \\__,_|_|\\_\\\033[0m")
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}

func SetConfig(config *config.Config) {
	cfg = config
}
