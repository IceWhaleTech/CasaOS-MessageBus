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

func (r *APIRoute) GetActionTypesBySourceID(ctx echo.Context, sourceID codegen.SourceID) error {
	panic("implement me") // TODO: Implement
}

func (r *APIRoute) GetActionType(ctx echo.Context, sourceID codegen.SourceID, name codegen.EventName) error {
	panic("implement me") // TODO: Implement
}

func (r *APIRoute) TriggerAction(ctx echo.Context, sourceID codegen.SourceID, name codegen.EventName) error {
	panic("implement me") // TODO: Implement
}

func (r *APIRoute) SubscribeAction(ctx echo.Context, sourceID codegen.SourceID, params codegen.SubscribeActionParams) error {
	panic("implement me") // TODO: Implement
}
