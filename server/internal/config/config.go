package config

import (
	"github.com/spf13/viper"
)

const DefaultListen = "localhost:8080"

type Config struct {
	Listen   string             `mapstructure:"listen"`
	Postgres PostgreSQLSettings `mapstructure:"postgres"`
}

type PostgreSQLSettings struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func NewConfig() *Config {
	return &Config{
		Listen:   DefaultListen,
		Postgres: PostgreSQLSettings{},
	}
}

func (c *Config) ReadFile(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}
