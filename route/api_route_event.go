package route

import (
	"net/http"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/route/adapter/in"
	"github.com/IceWhaleTech/CasaOS-MessageBus/route/adapter/out"
	"github.com/labstack/echo/v4"
)

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
	var event codegen.Event
	if err := ctx.Bind(&event); err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.ResponseBadRequest{Message: &message})
	}

	result, err := r.services.EventTypeService.Publish(in.EventAdapter(event))
	if err != nil {
		message := err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{Message: &message})
	}

	return ctx.JSON(http.StatusOK, out.EventAdapter(*result))
}
