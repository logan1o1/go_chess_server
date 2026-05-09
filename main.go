package main

import (
	_ "github.com/lib/pq"
	"github.com/logan1o1/go_chess_server/config"
	"github.com/logan1o1/go_chess_server/database"
)

func main() {
	config.LoadEnvVars()
	config.InitializeLogger()

	dbClient, err := database.NewPgClient(config.EnvVars.DbUser, config.EnvVars.DbPassword, config.EnvVars.DbName)
	if err != nil {
		config.ZapLogger.Panic("unable to connect database: " + err.Error())
	}

	app := InitRouter(dbClient)

	err = app.Run(":" + config.EnvVars.AppPort)
	if err != nil {
		config.ZapLogger.Panic("Unable to start the server: " + err.Error())
	}
}
