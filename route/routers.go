package route

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/service"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

func NewAPIRouter(swagger *openapi3.T, services service.Services) (http.Handler, error) {
	apiRoute := NewAPIRoute(services)

	e := echo.New()

	e.Use((echo_middleware.CORSWithConfig(echo_middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderXCSRFToken, echo.HeaderContentType, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowMethods, echo.HeaderConnection, echo.HeaderOrigin, echo.HeaderXRequestedWith},
		ExposeHeaders:    []string{echo.HeaderContentLength, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders},
		MaxAge:           172800,
		AllowCredentials: true,
	})))

	e.Use(echo_middleware.Gzip())

	e.Use(echo_middleware.Logger())

	e.Use(middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{Options: openapi3filter.Options{AuthenticationFunc: openapi3filter.NoopAuthenticationFunc}}))

	apiPath, err := getAPIPath(getSwaggerURL(swagger))
	if err != nil {
		return nil, err
	}

	codegen.RegisterHandlersWithBaseURL(e, apiRoute, apiPath)

	return e, nil
}

func NewDocRouter(swagger *openapi3.T, docHTML string, docYAML string) (http.Handler, error) {
	apiPath, err := getAPIPath(getSwaggerURL(swagger))
	if err != nil {
		return nil, err
	}

	docPath := "/doc" + apiPath

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == docPath {
			if _, err := w.Write([]byte(docHTML)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		if r.URL.Path == docPath+"/openapi.yaml" {
			if _, err := w.Write([]byte(docYAML)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}), nil
}

func NewWebSocketRouter(services service.Services) http.Handler {
	wsRoute := NewWebSocketRoute(services)

	e := echo.New()

	e.Use(echo_middleware.Gzip())

	e.Use(echo_middleware.Logger())

	v2Group := e.Group("/ws")
	v2WebSocketGroup := v2Group.Group("/v2")

	v2WebSocketMessageBusGroup := v2WebSocketGroup.Group("/message_bus")
	v2WebSocketMessageBusGroup.GET("/event_type/:source_id/:name", wsRoute.SubscribeEvents)
	v2WebSocketMessageBusGroup.GET("/action_type/:source_id/:name", wsRoute.SubscribeActions)

	return e
}

func getSwaggerURL(swagger *openapi3.T) string {
	return swagger.Servers[0].URL
}

func getAPIPath(swaggerURL string) (string, error) {
	u, err := url.Parse(swaggerURL)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(u.Path, "/"), nil
}
