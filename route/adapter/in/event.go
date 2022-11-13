package in

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
)

func EventAdapter(event codegen.Event) model.Event {
	properties := make([]model.Property, 0)
	for _, property := range event.Properties {
		properties = append(properties, PropertyAdapter(property))
	}

	var timestamp int64
	if event.Timestamp != nil {
		timestamp = event.Timestamp.Unix()
	}

	return model.Event{
		SourceID:   event.SourceID,
		Name:       event.Name,
		Properties: properties,
		Timestamp:  timestamp,
	}
}
