package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

var EnvVars EnvConfig

type EnvConfig struct {
	Environment string `env:"ENVIRONMENT"`
	AppPort     string `env:"PORT"`
	DbUser      string `env:"DB_USERNAME"`
	DbPassword  string `env:"DB_PASSWORD"`
	DbName      string `env:"DB_NAME"`
}

func LoadEnvVars() {
	godotenv.Load()
	if err := env.Parse(&EnvVars); err != nil {
		panic("Unable to load env vars: " + err.Error())
	}
}
