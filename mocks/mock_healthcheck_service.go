package mocks

import (
	"github.com/logan1o1/go_chess_server/interfaces"
	"github.com/logan1o1/go_chess_server/models"
)

type MockHealthCheckService struct {
	Response models.HealthCheckResponse
	Err      error
}

func (m *MockHealthCheckService) CheckHealthCheck() (models.HealthCheckResponse, error) {
	return m.Response, m.Err
}

var _ interfaces.IHealthCheckService = (*MockHealthCheckService)(nil)
