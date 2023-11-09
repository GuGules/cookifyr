package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	HttpAddress  string `env:"HTTP_ADDRESS" envDefault:":3000"`
	DbDriver     string `env:"DB_DRIVER" envDefault:"postgres"`
	DbConnString string `env:"DB_CONN_STRING"`
}

func NewFromEnv(files ...string) (Config, error) {
	config := Config{}

	if errLoad := godotenv.Load(files...); errLoad != nil {
		return Config{}, fmt.Errorf("config: %w", errLoad)
	}

	if errParse := env.Parse(&config); errParse != nil {
		return Config{}, fmt.Errorf("config: %w", errParse)
	}

	return config, nil
}
