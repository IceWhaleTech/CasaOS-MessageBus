// Package codegen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package codegen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	Access_tokenScopes = "access_token.Scopes"
)

// Defines values for YSKCardCardType.
const (
	YSKCardCardTypeLongNotice  YSKCardCardType = "long-notice"
	YSKCardCardTypeShortNotice YSKCardCardType = "short-notice"
	YSKCardCardTypeTask        YSKCardCardType = "task"
)

// Defines values for YSKCardRenderType.
const (
	YSKCardRenderTypeIconTextNotice YSKCardRenderType = "icon-text-notice"
	YSKCardRenderTypeListNotice     YSKCardRenderType = "list-notice"
	YSKCardRenderTypeMarkdownNotice YSKCardRenderType = "markdown-notice"
	YSKCardRenderTypeTask           YSKCardRenderType = "task"
)

// Action defines model for Action.
type Action struct {
	// Name action name
	Name string `json:"name"`

	// Properties event properties
	Properties map[string]string `json:"properties"`

	// SourceID associated source id
	SourceID string `json:"sourceID"`

	// Timestamp timestamp this action took place
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// ActionType defines model for ActionType.
type ActionType struct {
	// Name action name
	//
	// (there is no naming convention for action names, but it is recommended to name each as structural and descriptive as possible)
	Name             string         `json:"name"`
	PropertyTypeList []PropertyType `json:"propertyTypeList"`

	// SourceID action source id to identify where the action will take
	SourceID string `json:"sourceID"`
}

// BaseResponse defines model for BaseResponse.
type BaseResponse struct {
	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// Event defines model for Event.
type Event struct {
	// Name event name
	Name string `json:"name"`

	// Properties event properties
	Properties map[string]string `json:"properties"`

	// SourceID associated source id
	SourceID string `json:"sourceID"`

	// Timestamp timestamp this event took place
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// Uuid event uuid
	Uuid *string `json:"uuid,omitempty"`
}

// EventType defines model for EventType.
type EventType struct {
	// Name event name
	//
	// (there is no naming convention for event names, but it is recommended to name each as structural and descriptive as possible)
	Name             string         `json:"name"`
	PropertyTypeList []PropertyType `json:"propertyTypeList"`

	// SourceID event source id to identify where the event comes from
	SourceID string `json:"sourceID"`
}

// PropertyType defines model for PropertyType.
type PropertyType struct {
	Description *string `json:"description,omitempty"`
	Example     *string `json:"example,omitempty"`

	// Name property name
	//
	// > It is recommended for a property name to be as descriptive as possible. One option is to prefix with a namespace.
	// > - If the property is source specific, prefix with source ID. For example, `local-storage:vendor`
	// > - Otherwise, prefix with `common:`. For example, `common:email`
	// >
	// > Some bad examples are `id`, `avail`, `blk`...which can be ambiguous and confusing.
	Name string `json:"name"`
}

// YSKCard defines model for YSKCard.
type YSKCard struct {
	CardType   YSKCardCardType   `json:"cardType"`
	Content    YSKCardContent    `json:"content"`
	Id         string            `json:"id"`
	RenderType YSKCardRenderType `json:"renderType"`
}

// YSKCardCardType defines model for YSKCard.CardType.
type YSKCardCardType string

// YSKCardRenderType defines model for YSKCard.RenderType.
type YSKCardRenderType string

// YSKCardContent defines model for YSKCardContent.
type YSKCardContent struct {
	BodyIconWithText *YSKCardIconWithText   `json:"bodyIconWithText,omitempty"`
	BodyList         *[]YSKCardListItem     `json:"bodyList,omitempty"`
	BodyProgress     *YSKCardProgress       `json:"bodyProgress,omitempty"`
	FooterActions    *[]YSKCardFooterAction `json:"footerActions,omitempty"`
	TitleIcon        YSKCardIcon            `json:"titleIcon"`
	TitleText        string                 `json:"titleText"`
}

// YSKCardFooterAction defines model for YSKCardFooterAction.
type YSKCardFooterAction struct {
	MessageBus YSKCardMessageBusAction `json:"messageBus"`
	Side       string                  `json:"side"`
	Style      string                  `json:"style"`
	Text       string                  `json:"text"`
}

// YSKCardIcon defines model for YSKCardIcon.
type YSKCardIcon = string

// YSKCardIconWithText defines model for YSKCardIconWithText.
type YSKCardIconWithText struct {
	Description string      `json:"description"`
	Icon        YSKCardIcon `json:"icon"`
}

// YSKCardList defines model for YSKCardList.
type YSKCardList = []YSKCard

// YSKCardListItem defines model for YSKCardListItem.
type YSKCardListItem struct {
	Description string      `json:"description"`
	Icon        YSKCardIcon `json:"icon"`
	RightText   string      `json:"rightText"`
}

// YSKCardMessageBusAction defines model for YSKCardMessageBusAction.
type YSKCardMessageBusAction struct {
	Key     string `json:"key"`
	Payload string `json:"payload"`
}

// YSKCardProgress defines model for YSKCardProgress.
type YSKCardProgress struct {
	Label    string `json:"label"`
	Progress int    `json:"progress"`
}

// ActionName defines model for ActionName.
type ActionName = string

// ActionNames defines model for ActionNames.
type ActionNames = []string

// EventName defines model for EventName.
type EventName = string

// EventNames defines model for EventNames.
type EventNames = []string

// SourceID defines model for SourceID.
type SourceID = string

// GetActionTypeOK defines model for GetActionTypeOK.
type GetActionTypeOK = ActionType

// GetActionTypesOK defines model for GetActionTypesOK.
type GetActionTypesOK = []ActionType

// GetEventTypeOK defines model for GetEventTypeOK.
type GetEventTypeOK = EventType

// GetEventTypesOK defines model for GetEventTypesOK.
type GetEventTypesOK = []EventType

// PublishEventOK defines model for PublishEventOK.
type PublishEventOK = Event

// ResponseBadRequest defines model for ResponseBadRequest.
type ResponseBadRequest = BaseResponse

// ResponseConflict defines model for ResponseConflict.
type ResponseConflict = BaseResponse

// ResponseGetYSKCardListOK defines model for ResponseGetYSKCardListOK.
type ResponseGetYSKCardListOK struct {
	Data *YSKCardList `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// ResponseInternalServerError defines model for ResponseInternalServerError.
type ResponseInternalServerError = BaseResponse

// ResponseNotFound defines model for ResponseNotFound.
type ResponseNotFound = BaseResponse

// ResponseOK defines model for ResponseOK.
type ResponseOK = BaseResponse

// TriggerActionOK defines model for TriggerActionOK.
type TriggerActionOK = Action

// PublishEvent event properties
type PublishEvent map[string]string

// RegisterActionTypes defines model for RegisterActionTypes.
type RegisterActionTypes = []ActionType

// RegisterEventTypes defines model for RegisterEventTypes.
type RegisterEventTypes = []EventType

// TriggerAction action properties
type TriggerAction map[string]string

// SubscribeActionWSParams defines parameters for SubscribeActionWS.
type SubscribeActionWSParams struct {
	Names *ActionNames `form:"names,omitempty" json:"names,omitempty"`
}

// TriggerActionJSONBody defines parameters for TriggerAction.
type TriggerActionJSONBody map[string]string

// RegisterActionTypesJSONBody defines parameters for RegisterActionTypes.
type RegisterActionTypesJSONBody = []ActionType

// SubscribeEventWSParams defines parameters for SubscribeEventWS.
type SubscribeEventWSParams struct {
	Names *EventNames `form:"names,omitempty" json:"names,omitempty"`
}

// PublishEventJSONBody defines parameters for PublishEvent.
type PublishEventJSONBody map[string]string

// RegisterEventTypesJSONBody defines parameters for RegisterEventTypes.
type RegisterEventTypesJSONBody = []EventType

// TriggerActionJSONRequestBody defines body for TriggerAction for application/json ContentType.
type TriggerActionJSONRequestBody TriggerActionJSONBody

// RegisterActionTypesJSONRequestBody defines body for RegisterActionTypes for application/json ContentType.
type RegisterActionTypesJSONRequestBody = RegisterActionTypesJSONBody

// PublishEventJSONRequestBody defines body for PublishEvent for application/json ContentType.
type PublishEventJSONRequestBody PublishEventJSONBody

// RegisterEventTypesJSONRequestBody defines body for RegisterEventTypes for application/json ContentType.
type RegisterEventTypesJSONRequestBody = RegisterEventTypesJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Subscribe to actions by source ID (WebSocket)
	// (GET /action/{source_id})
	SubscribeActionWS(ctx echo.Context, sourceId SourceID, params SubscribeActionWSParams) error
	// Trigger an action
	// (POST /action/{source_id}/{name})
	TriggerAction(ctx echo.Context, sourceId SourceID, name ActionName) error
	// List action types
	// (GET /action_type)
	GetActionTypes(ctx echo.Context) error
	// Register one or more action types
	// (POST /action_type)
	RegisterActionTypes(ctx echo.Context) error
	// Get action types by source ID
	// (GET /action_type/{source_id})
	GetActionTypesBySourceID(ctx echo.Context, sourceId SourceID) error
	// Get an action type by source ID and name
	// (GET /action_type/{source_id}/{name})
	GetActionType(ctx echo.Context, sourceId SourceID, name ActionName) error
	// Subscribe to events by source ID (WebSocket)
	// (GET /event/{source_id})
	SubscribeEventWS(ctx echo.Context, sourceId SourceID, params SubscribeEventWSParams) error
	// Publish an event
	// (POST /event/{source_id}/{name})
	PublishEvent(ctx echo.Context, sourceId SourceID, name EventName) error
	// List event types
	// (GET /event_type)
	GetEventTypes(ctx echo.Context) error
	// Register one or more event types
	// (POST /event_type)
	RegisterEventTypes(ctx echo.Context) error
	// Get event types by source ID
	// (GET /event_type/{source_id})
	GetEventTypesBySourceID(ctx echo.Context, sourceId SourceID) error
	// Get an event type by source ID and name
	// (GET /event_type/{source_id}/{name})
	GetEventType(ctx echo.Context, sourceId SourceID, name EventName) error
	// Subscribe to events and actions (SocketIO)
	// (GET /socket.io)
	SubscribeSIO(ctx echo.Context) error
	// Poll events and actions (SocketIO)
	// (POST /socket.io)
	PollSIO(ctx echo.Context) error
	// Subscribe to events and actions (SocketIO)
	// (GET /socket.io/)
	SubscribeSIO2(ctx echo.Context) error
	// Poll events and actions (SocketIO)
	// (POST /socket.io/)
	PollSIO2(ctx echo.Context) error

	// (GET /ysk)
	GetYskCard(ctx echo.Context) error

	// (DELETE /ysk/{id})
	DeleteYskCard(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// SubscribeActionWS converts echo context to params.
func (w *ServerInterfaceWrapper) SubscribeActionWS(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source_id" -------------
	var sourceId SourceID

	err = runtime.BindStyledParameterWithLocation("simple", false, "source_id", runtime.ParamLocationPath, ctx.Param("source_id"), &sourceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source_id: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SubscribeActionWSParams
	// ------------- Optional query parameter "names" -------------

	err = runtime.BindQueryParameter("form", true, false, "names", ctx.QueryParams(), &params.Names)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter names: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SubscribeActionWS(ctx, sourceId, params)
	return err
}

// TriggerAction converts echo context to params.
func (w *ServerInterfaceWrapper) TriggerAction(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source_id" -------------
	var sourceId SourceID

	err = runtime.BindStyledParameterWithLocation("simple", false, "source_id", runtime.ParamLocationPath, ctx.Param("source_id"), &sourceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source_id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name ActionName

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TriggerAction(ctx, sourceId, name)
	return err
}

// GetActionTypes converts echo context to params.
func (w *ServerInterfaceWrapper) GetActionTypes(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetActionTypes(ctx)
	return err
}

// RegisterActionTypes converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterActionTypes(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterActionTypes(ctx)
	return err
}

// GetActionTypesBySourceID converts echo context to params.
func (w *ServerInterfaceWrapper) GetActionTypesBySourceID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source_id" -------------
	var sourceId SourceID

	err = runtime.BindStyledParameterWithLocation("simple", false, "source_id", runtime.ParamLocationPath, ctx.Param("source_id"), &sourceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source_id: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetActionTypesBySourceID(ctx, sourceId)
	return err
}

// GetActionType converts echo context to params.
func (w *ServerInterfaceWrapper) GetActionType(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source_id" -------------
	var sourceId SourceID

	err = runtime.BindStyledParameterWithLocation("simple", false, "source_id", runtime.ParamLocationPath, ctx.Param("source_id"), &sourceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source_id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name ActionName

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetActionType(ctx, sourceId, name)
	return err
}

// SubscribeEventWS converts echo context to params.
func (w *ServerInterfaceWrapper) SubscribeEventWS(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source_id" -------------
	var sourceId SourceID

	err = runtime.BindStyledParameterWithLocation("simple", false, "source_id", runtime.ParamLocationPath, ctx.Param("source_id"), &sourceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source_id: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SubscribeEventWSParams
	// ------------- Optional query parameter "names" -------------

	err = runtime.BindQueryParameter("form", true, false, "names", ctx.QueryParams(), &params.Names)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter names: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SubscribeEventWS(ctx, sourceId, params)
	return err
}

// PublishEvent converts echo context to params.
func (w *ServerInterfaceWrapper) PublishEvent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source_id" -------------
	var sourceId SourceID

	err = runtime.BindStyledParameterWithLocation("simple", false, "source_id", runtime.ParamLocationPath, ctx.Param("source_id"), &sourceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source_id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name EventName

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PublishEvent(ctx, sourceId, name)
	return err
}

// GetEventTypes converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventTypes(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetEventTypes(ctx)
	return err
}

// RegisterEventTypes converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterEventTypes(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterEventTypes(ctx)
	return err
}

// GetEventTypesBySourceID converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventTypesBySourceID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source_id" -------------
	var sourceId SourceID

	err = runtime.BindStyledParameterWithLocation("simple", false, "source_id", runtime.ParamLocationPath, ctx.Param("source_id"), &sourceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source_id: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetEventTypesBySourceID(ctx, sourceId)
	return err
}

// GetEventType converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventType(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "source_id" -------------
	var sourceId SourceID

	err = runtime.BindStyledParameterWithLocation("simple", false, "source_id", runtime.ParamLocationPath, ctx.Param("source_id"), &sourceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter source_id: %s", err))
	}

	// ------------- Path parameter "name" -------------
	var name EventName

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetEventType(ctx, sourceId, name)
	return err
}

// SubscribeSIO converts echo context to params.
func (w *ServerInterfaceWrapper) SubscribeSIO(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SubscribeSIO(ctx)
	return err
}

// PollSIO converts echo context to params.
func (w *ServerInterfaceWrapper) PollSIO(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PollSIO(ctx)
	return err
}

// SubscribeSIO2 converts echo context to params.
func (w *ServerInterfaceWrapper) SubscribeSIO2(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SubscribeSIO2(ctx)
	return err
}

// PollSIO2 converts echo context to params.
func (w *ServerInterfaceWrapper) PollSIO2(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PollSIO2(ctx)
	return err
}

// GetYskCard converts echo context to params.
func (w *ServerInterfaceWrapper) GetYskCard(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetYskCard(ctx)
	return err
}

// DeleteYskCard converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteYskCard(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteYskCard(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/action/:source_id", wrapper.SubscribeActionWS)
	router.POST(baseURL+"/action/:source_id/:name", wrapper.TriggerAction)
	router.GET(baseURL+"/action_type", wrapper.GetActionTypes)
	router.POST(baseURL+"/action_type", wrapper.RegisterActionTypes)
	router.GET(baseURL+"/action_type/:source_id", wrapper.GetActionTypesBySourceID)
	router.GET(baseURL+"/action_type/:source_id/:name", wrapper.GetActionType)
	router.GET(baseURL+"/event/:source_id", wrapper.SubscribeEventWS)
	router.POST(baseURL+"/event/:source_id/:name", wrapper.PublishEvent)
	router.GET(baseURL+"/event_type", wrapper.GetEventTypes)
	router.POST(baseURL+"/event_type", wrapper.RegisterEventTypes)
	router.GET(baseURL+"/event_type/:source_id", wrapper.GetEventTypesBySourceID)
	router.GET(baseURL+"/event_type/:source_id/:name", wrapper.GetEventType)
	router.GET(baseURL+"/socket.io", wrapper.SubscribeSIO)
	router.POST(baseURL+"/socket.io", wrapper.PollSIO)
	router.GET(baseURL+"/socket.io/", wrapper.SubscribeSIO2)
	router.POST(baseURL+"/socket.io/", wrapper.PollSIO2)
	router.GET(baseURL+"/ysk", wrapper.GetYskCard)
	router.DELETE(baseURL+"/ysk/:id", wrapper.DeleteYskCard)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"V+ur9f8DAAD//z3VnfHySgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
