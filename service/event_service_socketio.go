package service

import (
	"context"
	"net/http"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
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
		logger.Error("socketio serve error", zap.Error(err))
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

func buildServer() *socketio.Server {
	websocketTransport := websocket.Default
	websocketTransport.CheckOrigin = func(r *http.Request) bool {
		return true // TOOD remove this debug setting
	}

	pollingTransport := polling.Default
	pollingTransport.CheckOrigin = func(r *http.Request) bool {
		return true // TODO remove this debug setting
	}

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			websocketTransport,
			pollingTransport,
		},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		logger.Info("connected", zap.Any("id", s.ID()))
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		logger.Error("meet error", zap.Any("error", e))
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		logger.Info("closed", zap.Any("reason", reason))
	})

	return server
}
