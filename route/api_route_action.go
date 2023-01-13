package route

import (
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
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (r *APIRoute) GetActionTypes(c echo.Context) error {
	actionType, err := r.services.ActionTypeService.GetActionTypes()
	if err != nil {
		message := err.Error()
		return c.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	results := make([]codegen.ActionType, 0)

	for _, actionType := range actionType {
		results = append(results, out.ActionTypeAdapter(actionType))
	}

	return c.JSON(http.StatusOK, results)
}

func (r *APIRoute) RegisterActionTypes(c echo.Context) error {
	var actionTypes []codegen.ActionType
	if err := c.Bind(&actionTypes); err != nil {
		message := err.Error()
		return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	for _, actionType := range actionTypes {
		_, err := r.services.ActionTypeService.RegisterActionType(in.ActionTypeAdapter(actionType))
		if err != nil {
			message := err.Error()
			return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
		}
	}

	return c.JSON(http.StatusOK, codegen.ResponseOK{})
}

func (r *APIRoute) GetActionTypesBySourceID(c echo.Context, sourceID codegen.SourceID) error {
	results, err := r.services.ActionTypeService.GetActionTypesBySourceID(sourceID)
	if err != nil {
		message := err.Error()
		return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	return c.JSON(http.StatusOK, results)
}

func (r *APIRoute) GetActionType(c echo.Context, sourceID codegen.SourceID, name codegen.EventName) error {
	result, err := r.services.ActionTypeService.GetActionType(sourceID, name)
	if err != nil {
		message := err.Error()
		return c.JSON(http.StatusNotFound, codegen.ResponseNotFound{Message: &message})
	}

	if result == nil {
		return c.JSON(http.StatusNotFound, codegen.ResponseNotFound{Message: utils.Ptr("not found")})
	}

	return c.JSON(http.StatusOK, result)
}

func (r *APIRoute) TriggerAction(c echo.Context, sourceID codegen.SourceID, name codegen.EventName) error {
	actionType, err := r.services.ActionTypeService.GetActionType(sourceID, name)
	if err != nil {
		message := err.Error()
		return c.JSON(http.StatusNotFound, codegen.ResponseNotFound{Message: &message})
	}

	if actionType == nil {
		return c.JSON(http.StatusNotFound, codegen.ResponseNotFound{Message: utils.Ptr("not found")})
	}

	var properties map[string]string
	if err := c.Bind(&properties); err != nil {
		message := err.Error()
		return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	action := in.ActionAdapter(codegen.Action{
		SourceID:   sourceID,
		Name:       name,
		Properties: properties,
		Timestamp:  utils.Ptr(time.Now()),
	})

	go r.services.SocketIOService.Publish(action)
	go r.services.ActionServiceWS.Trigger(action)

	return c.JSON(http.StatusOK, out.ActionAdapter(action))
}

func (r *APIRoute) SubscribeActionWS(c echo.Context, sourceID codegen.SourceID, params codegen.SubscribeActionWSParams) error {
	var actionNames []string
	if params.Names != nil {
		for _, actionName := range *params.Names {
			actionType, err := r.services.ActionTypeService.GetActionType(sourceID, actionName)
			if err != nil {
				message := err.Error()
				return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
			}

			if actionType == nil {
				return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: utils.Ptr("action type not found")})
			}

			actionNames = append(actionNames, actionName)
		}
	} else {
		actionTypes, err := r.services.ActionTypeService.GetActionTypesBySourceID(sourceID)
		if err != nil {
			message := err.Error()
			return c.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
		}

		for _, actionType := range actionTypes {
			actionNames = append(actionNames, actionType.Name)
		}
	}

	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
	if err != nil {
		message := err.Error()
		return c.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	channel, err := r.services.ActionServiceWS.Subscribe(sourceID, actionNames)
	if err != nil {
		conn.Close() // need to close connection here, instead of defer, because of the goroutine
		message := err.Error()
		return c.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	go func(conn net.Conn, channel chan model.Action, actionNames []string) {
		defer conn.Close()
		defer close(channel)
		defer func(actionNames []string) {
			for _, name := range actionNames {
				if err := r.services.ActionServiceWS.Unsubscribe(sourceID, name, channel); err != nil {
					logger.Error("error when trying to unsubscribe an action type via websocket", zap.Error(err), zap.String("source_id", sourceID), zap.String("name", name))
				}
			}
		}(actionNames)

		logger.Info("a websocket connection has started for actions", zap.String("remote_addr", conn.RemoteAddr().String()))

		for {
			action, ok := <-channel
			if !ok {
				logger.Info("websocket channel for events is closed")
				return
			}

			if action.SourceID == common.MessageBusSourceID && action.Name == common.MessageBusHeartbeatName {
				if err := wsutil.WriteServerMessage(conn, ws.OpPing, []byte{}); err != nil {
					logger.Error("error when trying to send ping message via websocket", zap.Error(err))
					return
				}
				continue
			}

			message, err := json.Marshal(out.ActionAdapter(action))
			if err != nil {
				logger.Error("error when trying to marshal action for websocket", zap.Error(err))
				continue
			}

			logger.Info("sending action via websocket", zap.String("remote_addr", conn.RemoteAddr().String()), zap.String("message", string(message)))

			if err := wsutil.WriteServerBinary(conn, message); err != nil {
				if _, ok := err.(*net.OpError); ok {
					logger.Info("websocket connection ended", zap.String("error", err.Error()))
				} else {
					logger.Error("error when sending event via websocket", zap.String("error", err.Error()))
				}
				return
			}
		}
	}(conn, channel, actionNames)

	return nil
}
