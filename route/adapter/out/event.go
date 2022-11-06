package out

import (
	"time"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
)

func EventAdapter(event model.Event) codegen.Event {
	properties := make([]codegen.Property, 0)
	for _, property := range event.Properties {
		properties = append(properties, PropertyAdapter(property))
	}

	timestamp := time.Unix(event.Timestamp, 0)

	return codegen.Event{
		SourceID:   &event.SourceID,
		Name:       &event.Name,
		Properties: &properties,
		Timestamp:  &timestamp,
	}
}
