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

func NewConfig(listen, pg_host, pg_port, user, password, db string) *Config {
	if listen == "" {
		listen = DefaultListen
	}
	return &Config{
		Listen: listen,
		Postgres: PostgreSQLSettings{
			Host:     pg_host,
			Port:     pg_port,
			User:     user,
			Password: password,
			Database: db,
		},
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
