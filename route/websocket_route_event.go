package route

import (
	"net"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/route/adapter/out"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (r *WebSocketRoute) SubscribeEvents(c echo.Context) error {
	sourceID := c.Param("source_id")
	name := c.Param("name")

	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
	if err != nil {
		return err
	}

	channel, err := r.services.EventTypeService.Subscribe(sourceID, name)
	if err != nil {
		conn.Close() // need to close connection here, instead of defer, because of the goroutine
		return err
	}

	go func(conn net.Conn, channel chan model.Event) {
		defer conn.Close()
		defer func() {
			if err := r.services.EventTypeService.Unsubscribe(sourceID, name, channel); err != nil {
				logger.Error("error when trying to unsubscribe an event type", zap.Error(err), zap.String("source_id", sourceID), zap.String("name", name))
			}
		}()

		logger.Info("started", zap.String("remote_addr", conn.RemoteAddr().String()))

		for {
			event, ok := <-channel
			if !ok {
				logger.Info("channel closed")
				return
			}

			message, err := json.Marshal(out.EventAdapter(event))
			if err != nil {
				logger.Error("failed to marshal event", zap.Error(err))
				continue
			}

			logger.Info("sending", zap.String("remote_addr", conn.RemoteAddr().String()), zap.String("message", string(message)))

			if err := wsutil.WriteServerMessage(conn, ws.OpText, message); err != nil {
				if _, ok := err.(*net.OpError); ok {
					logger.Info("ended", zap.String("error", err.Error()))
				} else {
					logger.Error("error", zap.String("error", err.Error()))
				}
				return
			}
		}
	}(conn, channel)

	return nil
}
