package main

import (
	"sync"

	"github.com/logan1o1/go_chess_server/controllers/healthcheck"
	"github.com/logan1o1/go_chess_server/interfaces"
	"github.com/logan1o1/go_chess_server/services/healthcheck_service"
)

type ServiceInjector struct{}

func (s *ServiceInjector) InjectHealthCheckController(db interfaces.IDatabase) *healthcheck.HealthCheckController {
	return healthcheck.NewHealthCkeckController(healthcheck_service.NewHealthCheckService(db))
}

var (
	serviceInjectorObj *ServiceInjector
	injectOnce         sync.Once
)

func NewServiceInjector() *ServiceInjector {
	injectOnce.Do(func() {
		serviceInjectorObj = &ServiceInjector{}
	})
	return serviceInjectorObj
}
