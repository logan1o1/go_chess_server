package main

import (
	"github.com/gin-gonic/gin"
	"github.com/logan1o1/go_chess_server/interfaces"
)

func InitRouter(db interfaces.IDatabase) *gin.Engine {
	app := gin.Default()

	healthcheckController := NewServiceInjector().InjectHealthCheckController(db)
	app.GET("/healthcheck", healthcheckController.CheckServerHealthCheck)

	return app
}
