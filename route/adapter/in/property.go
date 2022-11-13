package in

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
)

func PropertyAdapter(property codegen.Property) model.Property {
	return model.Property{
		Name:  property.Name,
		Value: property.Value,
	}
}
