package service

import (
	"context"
	"errors"

	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type Services struct {
	EventService  *EventService
	ActionService *ActionService
}

var (
	ErrInboundChannelNotFound     = errors.New("inbound channel not found")
	ErrSubscriberChannelsNotFound = errors.New("subscriber channels not found")
)

func (s *Services) Start(ctx *context.Context) {
	go s.EventService.Start(ctx)
	go s.ActionService.Start(ctx)
}

func NewServices(repository *repository.Repository) Services {
	return Services{
		EventService:  NewEventService(repository),
		ActionService: NewActionService(repository),
	}
}
