package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"go.uber.org/zap"
)

type Services struct {
	EventTypeService *EventTypeService
	EventServiceWS   *EventServiceWS
	EventServiceSIO  *EventServiceSIO

	ActionTypeService *ActionTypeService
	ActionServiceWS   *ActionServiceWS
	ActionServiceSIO  *ActionServiceSIO
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
	go s.ActionServiceSIO.Start(ctx)
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
		ActionServiceSIO:  NewActionServiceSIO(),
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
