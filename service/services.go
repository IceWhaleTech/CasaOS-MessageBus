package service

import (
	"context"
	"errors"

	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type Services struct {
	EventTypeService *EventTypeService
	EventServiceWS   *EventServiceWS
	EventServiceSIO  *EventServiceSIO

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

	go s.EventServiceSIO.Start(ctx)
	// TODO - action SIO
}

func NewServices(repository *repository.Repository) Services {
	eventTypeService := NewEventTypeService(repository)
	actionTypeService := NewActionTypeService(repository)

	return Services{
		EventTypeService: eventTypeService,
		EventServiceWS:   NewEventServiceWS(eventTypeService),
		EventServiceSIO:  NewEventServiceSIO(),

		ActionTypeService: actionTypeService,
		ActionServiceWS:   NewActionServiceWS(actionTypeService),
	}
}
