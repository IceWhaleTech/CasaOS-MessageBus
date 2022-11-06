package service

import (
	"context"

	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type Services struct {
	EventTypeService EventTypeService
}

func (s *Services) Start() {
	go s.EventTypeService.Start()
}

func NewServices(ctx *context.Context, repository repository.Repository) Services {
	return Services{
		EventTypeService: NewEventTypeService(ctx, repository),
	}
}
