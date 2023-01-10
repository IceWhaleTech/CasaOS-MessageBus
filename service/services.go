package service

import (
	"context"
	"errors"

	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type Services struct {
	EventTypeService *EventTypeService
	EventServiceWS   *EventServiceWS

	ActionTypeService *ActionTypeService
	ActionServiceWS   *ActionServiceWS
}

var (
	ErrInboundChannelNotFound     = errors.New("inbound channel not found")
	ErrSubscriberChannelsNotFound = errors.New("subscriber channels not found")
	ErrAlreadySubscribed          = errors.New("already subscribed")
)

func (s *Services) Start(ctx *context.Context) {
	go s.EventServiceWS.Start(ctx)
	go s.ActionServiceWS.Start(ctx)
}

func NewServices(repository *repository.Repository) Services {
	return Services{
		EventTypeService:  NewEventTypeService(repository),
		ActionTypeService: NewActionTypeService(repository),
	}
}
