package in

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
)

func ActionAdapter(action codegen.Action) model.Action {
	var timestamp int64
	if action.Timestamp != nil {
		timestamp = action.Timestamp.Unix()
	}

	return model.Action{
		SourceID:   action.SourceID,
		Name:       action.Name,
		Properties: action.Properties,
		Timestamp:  timestamp,
	}
}
