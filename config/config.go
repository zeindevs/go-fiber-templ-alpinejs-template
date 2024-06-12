package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
}

func NewConfig() Config {
	return Config{}
}

func (c *Config) Get(key string) string {
	return os.Getenv(key)
}
