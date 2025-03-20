package main

import (
	"fmt"
	"os"

	"github.com/patinhooh/anitrak/cmd"
	"github.com/patinhooh/anitrak/internal/config"
)

func main() {
	// Load config once
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cmd.SetConfig(cfg)

	// Execute Cobra CLI
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
