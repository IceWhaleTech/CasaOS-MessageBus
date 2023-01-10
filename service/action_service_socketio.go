package service

import (
	"context"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	socketio "github.com/googollee/go-socket.io"
	"go.uber.org/zap"
)

type ActionServiceSIO struct {
	server *socketio.Server
}

func (s *ActionServiceSIO) Trigger(action model.Action) {
	s.server.BroadcastToNamespace("/", action.SourceID, action)
}

func (s *ActionServiceSIO) Start(ctx *context.Context) {
	if err := s.server.Serve(); err != nil {
		logger.Error("error when serving socketio for actions", zap.Error(err))
	}
}

func (s *ActionServiceSIO) Server() *socketio.Server {
	return s.server
}

func NewActionServiceSIO() *ActionServiceSIO {
	return &ActionServiceSIO{
		server: buildServer(),
	}
}
