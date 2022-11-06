package route

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/labstack/echo/v4"
)

func (r *APIRoute) GetActionTypes(ctx echo.Context) error {
	panic("implement me") // TODO: Implement
}

func (r *APIRoute) RegisterActionType(ctx echo.Context) error {
	panic("implement me") // TODO: Implement
}

func (r *APIRoute) GetActionTypesBySourceID(ctx echo.Context, sourceID codegen.SourceId) error {
	panic("implement me") // TODO: Implement
}

func (r *APIRoute) GetActionType(ctx echo.Context, sourceID codegen.SourceId, name codegen.Name) error {
	panic("implement me") // TODO: Implement
}

func (r *APIRoute) TriggerAction(ctx echo.Context, sourceID codegen.SourceId, name codegen.Name) error {
	panic("implement me") // TODO: Implement
}

func (r *WebSocketRoute) SubscribeActions(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
