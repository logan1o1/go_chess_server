package healthcheck_service

import (
	"github.com/logan1o1/go_chess_server/interfaces"
	"github.com/logan1o1/go_chess_server/models"
)

type HealthCheckService struct {
	db interfaces.IDatabase
}

func (hcs *HealthCheckService) CheckHealthCheck() (models.HealthCheckResponse, error) {
	return models.HealthCheckResponse{
		Status: "ok",
	}, nil
}

func NewHealthCheckService(db interfaces.IDatabase) *HealthCheckService {
	return &HealthCheckService{db: db}
}
