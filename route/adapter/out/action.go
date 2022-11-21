package out

import (
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils"
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
)

func ActionAdapter(action model.Action) codegen.Action {
	properties := make(codegen.Property)
	// for k, v := range action.Properties {
	// 	properties = append(properties, codegen.Property{Name: k, Value: v})
	// }
	properties = action.Properties

	return codegen.Action{
		SourceID:   &action.SourceID,
		Name:       &action.Name,
		Properties: &properties,
		Timestamp:  utils.Ptr(time.Unix(action.Timestamp, 0)),
	}
}
