package config

import "github.com/spf13/viper"

const DefaultListen = "localhost:8080"

type Config struct {
	Listen string `mapstructure:"listen"`
}

func NewConfig(listen string) *Config {
	if listen == "" {
		listen = DefaultListen
	}
	return &Config{
		Listen: listen,
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
