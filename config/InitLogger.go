package config

import "go.uber.org/zap"

var ZapLogger *zap.Logger

func InitializeLogger() {
	var err error
	if EnvVars.Environment == "DEV" {
		ZapLogger, err = zap.NewDevelopment()
		if err != nil {
			panic("Unable to intitialize logger")
		}
	}

	ZapLogger, err = zap.NewProduction()
	if err != nil {
		panic("Unable to initialize logger")
	}
}
