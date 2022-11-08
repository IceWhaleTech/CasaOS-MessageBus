package service

import (
	"context"

	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type Services struct {
	EventTypeService *EventTypeService
}

func (s *Services) Start(ctx *context.Context) {
	go s.EventTypeService.Start(ctx)
}

func NewServices(repository *repository.Repository) Services {
	return Services{
		EventTypeService: NewEventTypeService(repository),
	}
}
