package route

import "github.com/IceWhaleTech/CasaOS-MessageBus/service"

type WebSocketRoute struct {
	services service.Services
}

func NewWebSocketRoute(services service.Services) *WebSocketRoute {
	return &WebSocketRoute{
		services: services,
	}
}
