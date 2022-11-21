package in

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
)

func ActionAdapter(action codegen.Action) model.Action {
	properties := make(map[string]string)
	for _, property := range *action.Properties {
		properties[property.Name] = property.Value
	}

	var timestamp int64
	if action.Timestamp != nil {
		timestamp = action.Timestamp.Unix()
	}

	return model.Action{
		SourceID:   *action.SourceID,
		Name:       *action.Name,
		Properties: properties,
		Timestamp:  timestamp,
	}
}
