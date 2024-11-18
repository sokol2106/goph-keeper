package main

import (
	"fmt"
	"os"
	"server/internal/app"
	"server/internal/config"
)

const DefaultFileConfig = "config.json"

func main() {
	cfg := config.NewConfig(
		os.Getenv("LISTEN"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASE"),
	)

	fileConfig := os.Getenv("CONFIG")
	if fileConfig == "" {
		fileConfig = DefaultFileConfig
	}
	ParseFlags(cfg)
	err := cfg.ReadFile(fileConfig)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}

	app.Run(cfg)
}
