package main

import (
	"github.com/patinhooh/anitrak/cmd"
	"github.com/patinhooh/anitrak/internal/config"
)

func main() {
	config.InitConfig()
	cmd.Execute()
}
