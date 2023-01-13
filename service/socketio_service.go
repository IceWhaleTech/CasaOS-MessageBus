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

type SocketIOService struct {
	server *socketio.Server
}

func (s *SocketIOService) Publish(message interface{}) {
	if event, ok := message.(model.Event); ok {
		s.server.BroadcastToRoom("/", "event", event.Name, event)
		return
	}

	if action, ok := message.(model.Action); ok {
		s.server.BroadcastToRoom("/", "action", action.Name, action)
		return
	}

	logger.Error("unknown message type", zap.Any("message", message))
}

func (s *SocketIOService) Start(ctx *context.Context) {
	if err := s.server.Serve(); err != nil {
		logger.Error("error when serving socketio for events", zap.Error(err))
	}
}

func (s *SocketIOService) Server() *socketio.Server {
	return s.server
}

func NewSocketIOService() *SocketIOService {
	return &SocketIOService{
		server: buildServer(),
	}
}

func buildServer() *socketio.Server {
	websocketTransport := websocket.Default
	websocketTransport.CheckOrigin = func(r *http.Request) bool {
		return true // TODO remove this debug setting
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
		logger.Info("a socketio connection has started", zap.Any("remote_addr", s.RemoteAddr()))

		s.Join("event")
		s.Join("action")

		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		logger.Error("error in socketio connnection", zap.Any("error", e))
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		logger.Info("a socketio connection is disconnected", zap.Any("reason", reason))
	})

	return server
}
