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

	SocketIOService *SocketIOService
}

var (
	ErrInboundChannelNotFound     = errors.New("inbound channel not found")
	ErrSubscriberChannelsNotFound = errors.New("subscriber channels not found")
	ErrAlreadySubscribed          = errors.New("already subscribed")
)

func (s *Services) Start(ctx *context.Context) {
	go s.EventServiceWS.Start(ctx)
	go s.ActionServiceWS.Start(ctx)

	go s.SocketIOService.Start(ctx)
}

func NewServices(repository *repository.Repository) Services {
	eventTypeService := NewEventTypeService(repository)
	actionTypeService := NewActionTypeService(repository)

	return Services{
		EventTypeService: eventTypeService,
		EventServiceWS:   NewEventServiceWS(eventTypeService),
		SocketIOService:  NewSocketIOService(),

		ActionTypeService: actionTypeService,
		ActionServiceWS:   NewActionServiceWS(actionTypeService),
	}
}
