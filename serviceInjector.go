package main

import (
	"sync"
)

type ServiceInjector struct{}

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
