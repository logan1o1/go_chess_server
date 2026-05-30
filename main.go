package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/logan1o1/go_chess_server/config"
	"github.com/logan1o1/go_chess_server/database"
)

func main() {
	config.LoadEnvVars()

	dbClient, err := database.NewPgClient(
		config.EnvVars.DbUser,
		config.EnvVars.DbPassword,
		config.EnvVars.DbName)
	if err != nil {
		log.Panic("Unable to connect to database" + err.Error())
	}

	app := InitSshServer(dbClient)

	err = app.ListenAndServe()
	if err != nil {
		log.Panic("Unable to start the server: " + err.Error())
	}
}
