package route

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/common"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/route/adapter/in"
	"github.com/IceWhaleTech/CasaOS-MessageBus/route/adapter/out"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (r *APIRoute) GetEventTypes(ctx echo.Context) error {
	eventTypes, err := r.services.EventService.GetEventTypes()
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

	result, err := r.services.EventService.RegisterEventType(in.EventTypeAdapter(eventType))
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	return ctx.JSON(http.StatusOK, result)
}

func (r *APIRoute) GetEventTypesBySourceID(ctx echo.Context, sourceID codegen.SourceID) error {
	results, err := r.services.EventService.GetEventTypesBySourceID(sourceID)
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	return ctx.JSON(http.StatusOK, results)
}

func (r *APIRoute) GetEventType(ctx echo.Context, sourceID codegen.SourceID, name codegen.EventName) error {
	result, err := r.services.EventService.GetEventType(sourceID, name)
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusNotFound, codegen.ResponseNotFound{Message: &message})
	}

	if result == nil {
		return ctx.JSON(http.StatusNotFound, codegen.ResponseNotFound{Message: utils.Ptr("not found")})
	}

	return ctx.JSON(http.StatusOK, result)
}

func (r *APIRoute) PublishEvent(ctx echo.Context, sourceID codegen.SourceID, name codegen.EventName) error {
	eventType, err := r.services.EventService.GetEventType(sourceID, name)
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusNotFound, codegen.ResponseNotFound{Message: &message})
	}

	if eventType == nil {
		return ctx.JSON(http.StatusNotFound, codegen.ResponseNotFound{Message: utils.Ptr("not found")})
	}

	var properties map[string]string
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {

		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	if err = json.Unmarshal(body, &properties); err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	uuidStr := uuid.New().String()
	event := codegen.Event{
		SourceID:   sourceID,
		Name:       name,
		Properties: properties,
		Timestamp:  utils.Ptr(time.Now()),
		Uuid:       &uuidStr,
	}

	result, err := r.services.EventService.Publish(in.EventAdapter(event))
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	return ctx.JSON(http.StatusOK, out.EventAdapter(*result))
}

func (r *APIRoute) SubscribeEvent(c echo.Context, sourceID codegen.SourceID, params codegen.SubscribeEventParams) error {
	var eventNames []string
	if params.Names != nil {
		for _, eventName := range *params.Names {
			eventType, err := r.services.EventService.GetEventType(sourceID, eventName)
			if err != nil {
				message := err.Error()
				return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
			}

			if eventType == nil {
				return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: utils.Ptr(fmt.Sprintf("event type `%s` of source ID `%s` not found", eventName, sourceID))})
			}

			eventNames = append(eventNames, eventName)
		}
	} else {
		eventTypes, err := r.services.EventService.GetEventTypesBySourceID(sourceID)
		if err != nil || len(eventTypes) == 0 {
			message := err.Error()
			return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
		}

		for _, eventType := range eventTypes {
			eventNames = append(eventNames, eventType.Name)
		}
	}

	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
	if err != nil {
		message := err.Error()
		return c.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	channel, err := r.services.EventService.Subscribe(sourceID, eventNames)
	if err != nil {
		conn.Close() // need to close connection here, instead of defer, because of the goroutine
		message := err.Error()
		return c.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	go func(conn net.Conn, channel chan model.Event, eventNames []string) {
		defer conn.Close()
		defer close(channel)
		defer func(eventNames []string) {
			for _, name := range eventNames {
				if err := r.services.EventService.Unsubscribe(sourceID, name, channel); err != nil {
					logger.Error("error when trying to unsubscribe an event type", zap.Error(err), zap.String("source_id", sourceID), zap.String("name", name))
				}
			}
		}(eventNames)

		logger.Info("started", zap.String("remote_addr", conn.RemoteAddr().String()))

		for {
			event, ok := <-channel
			if !ok {
				logger.Info("channel closed")
				return
			}

			if event.SourceID == common.MessageBusSourceID && event.Name == common.MessageBusHeartbeatName {
				if err := wsutil.WriteServerMessage(conn, ws.OpPing, []byte{}); err != nil {
					logger.Error("error when trying to send ping message", zap.Error(err))
					return
				}
				continue
			}

			message, err := json.Marshal(out.EventAdapter(event))
			if err != nil {
				logger.Error("error when trying to marshal event", zap.Error(err))
				continue
			}

			logger.Info("sending", zap.String("remote_addr", conn.RemoteAddr().String()), zap.String("message", string(message)))

			if err := wsutil.WriteServerText(conn, message); err != nil {
				if _, ok := err.(*net.OpError); ok {
					logger.Info("ended", zap.String("error", err.Error()))
				} else {
					logger.Error("error", zap.String("error", err.Error()))
				}
				return
			}
		}
	}(conn, channel, eventNames)

	return nil
}
