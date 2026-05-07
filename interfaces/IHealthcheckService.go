package interfaces

import "github.com/logan1o1/go_chess_server/models"

type IHealthCheckService interface {
	CheckHealthCheck() (models.HealthCheckResponse, error)
}
