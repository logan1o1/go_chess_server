package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

var EnvVars EnvConfig

type EnvConfig struct {
	Environment string `env:"ENVIRONMENT"`
	AppPort     string `env:"PORT"`
	DbConnUrl   string `env:"POSTGRES_SQL_CONNECTION_STRING"`
}

func LoadEnvVars() {
	godotenv.Load()
	if err := env.Parse(&EnvVars); err != nil {
		panic("Unable to load env vars: " + err.Error())
	}
}
