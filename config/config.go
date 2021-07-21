package config

import (
	"fmt"

	"github.com/MrWebUzb/goenv"
)

// Config configuration values loaded from .env
// Environment equals one of the below values
// development, staging, production
// LogLevel equals one of the below values
// debug, info, warn, error, dpanic, panic, fatal
type Config struct {
	App              string `env:"APP" default:"example_app"`
	Environment      string `env:"ENVIRONMENT" default:"development"`
	LogLevel         string `env:"LOG_LEVEL" default:"debug"`
	HTTPPort         string `env:"HTTP_PORT" default:":8080"`
	DefaultOffset    string `env:"DEFAULT_OFFSET" default:"0"`
	DefaultLimit     string `env:"DEFAULT_LIMIT" default:"20"`
	PostgresHost     string `env:"POSTGRES_HOST" default:"localhost"`
	PostgresPort     int    `env:"POSTGRES_PORT" default:"5432"`
	PostgresDatabase string `env:"POSTGRES_DATABASE" default:"test"`
	PostgresUser     string `env:"POSTGRES_USER" default:"test"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" default:"test"`
}

// Load ...
func Load() Config {
	env, err := goenv.New()

	if err != nil {
		fmt.Printf("could not load env variables: %v\n", err)
		return Config{}
	}

	config := Config{}

	if err := env.Parse(&config); err != nil {
		fmt.Printf("could not parse env variables: %v\n", err)
		return Config{}
	}

	return config
}
