package main

import (
	"fmt"

	"github.com/patinhooh/anitrak/internal/config"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Println("Warning: Using default config due to error:", err)
	}

	fmt.Print(cfg)
}
