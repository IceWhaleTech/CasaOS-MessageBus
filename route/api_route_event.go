package route

import (
	"net"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/route/adapter/in"
	"github.com/IceWhaleTech/CasaOS-MessageBus/route/adapter/out"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (r *APIRoute) GetEventTypes(ctx echo.Context) error {
	eventTypes, err := r.services.EventTypeService.GetEventTypes()
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	results := make([]codegen.EventType, 0)

	for _, eventType := range eventTypes {
		results = append(results, out.EventTypeAdapter(eventType))
	}

	return ctx.JSON(http.StatusOK, results)
}

func (r *APIRoute) RegisterEventType(ctx echo.Context) error {
	var eventType codegen.EventType
	if err := ctx.Bind(&eventType); err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	result, err := r.services.EventTypeService.RegisterEventType(in.EventTypeAdapter(eventType))
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	return ctx.JSON(http.StatusOK, result)
}

func (r *APIRoute) GetEventTypesBySourceID(ctx echo.Context, sourceID codegen.SourceId) error {
	results, err := r.services.EventTypeService.GetEventTypesBySourceID(sourceID)
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	return ctx.JSON(http.StatusOK, results)
}

func (r *APIRoute) GetEventType(ctx echo.Context, sourceID codegen.SourceId, name codegen.Name) error {
	result, err := r.services.EventTypeService.GetEventType(sourceID, name)
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	return ctx.JSON(http.StatusOK, result)
}

func (r *APIRoute) PublishEvent(ctx echo.Context, sourceID codegen.SourceId, name codegen.Name) error {
	var properties []codegen.Property
	if err := ctx.Bind(&properties); err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	timestamp := time.Now()

	event := codegen.Event{
		SourceID:   &sourceID,
		Name:       &name,
		Properties: &properties,
		Timestamp:  &timestamp,
	}

	result, err := r.services.EventTypeService.Publish(in.EventAdapter(event))
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	return ctx.JSON(http.StatusOK, out.EventAdapter(*result))
}

func (r *APIRoute) SubscribeEvent(c echo.Context, sourceID codegen.SourceId, name codegen.Name) error {
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
