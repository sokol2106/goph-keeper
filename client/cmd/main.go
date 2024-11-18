package main

import (
	"client/internal/app"
	"client/internal/config"
	"fmt"
	"os"
)

const DefaultFileConfig = "config.json"

func main() {
	cfg := config.NewConfig("")
	err := cfg.ReadFile(DefaultFileConfig)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
	app.Run(cfg)
}
