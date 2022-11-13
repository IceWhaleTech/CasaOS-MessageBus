package out

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
)

func PropertyAdapter(property model.Property) codegen.Property {
	return codegen.Property{
		Name:  property.Name,
		Value: property.Value,
	}
}
