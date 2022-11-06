package out

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
)

func EventAdapter(event model.Event) codegen.Event {
	properties := make([]codegen.Property, 0)
	for _, property := range event.Properties {
		properties = append(properties, PropertyAdapter(property))
	}

	return codegen.Event{
		SourceID:   &event.SourceID,
		Name:       &event.Name,
		Properties: &properties,
	}
}
