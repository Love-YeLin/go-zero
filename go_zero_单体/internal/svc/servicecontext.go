package svc

import (
	"mail/internal/config"
	"mail/uc"
)

type ServiceContext struct {
	Config config.Config
	UC     *uc.UseCase
}

func NewServiceContext(c config.Config) *ServiceContext {
	useCase, _ := uc.NewUseCase(&c)
	return &ServiceContext{
		Config: c,
		UC:     useCase,
	}
}
