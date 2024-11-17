package main

import (
	"fmt"
	"os"
	"server/internal/app"
	"server/internal/config"
)

func main() {
	cfg := config.NewConfig()
	err := cfg.ReadFile("config.json")
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}

	app.Run(cfg)
}
