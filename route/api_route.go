package route

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/service"
)

type APIRoute struct {
	services *service.Services
}

func NewAPIRoute(services *service.Services) codegen.ServerInterface {
	return &APIRoute{
		services: services,
	}
}
