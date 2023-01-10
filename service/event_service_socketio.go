package service

import (
	"context"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	socketio "github.com/googollee/go-socket.io"
	"go.uber.org/zap"
)

type EventServiceSIO struct {
	server *socketio.Server
}

func (s *EventServiceSIO) Publish(event model.Event) {
	s.server.BroadcastToNamespace("/", event.SourceID, event)
}

func (s *EventServiceSIO) Start(ctx *context.Context) {
	if err := s.server.Serve(); err != nil {
		logger.Error("error when serving socketio for events", zap.Error(err))
	}
}

func (s *EventServiceSIO) Server() *socketio.Server {
	return s.server
}

func NewEventServiceSIO() *EventServiceSIO {
	return &EventServiceSIO{
		server: buildServer(),
	}
}
