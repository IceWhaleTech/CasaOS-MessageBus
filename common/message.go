package common

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/samber/lo"
)

const (
	SERVICENAME = "ysk"
)

// common properties
var (
	PropertyTypeMessage = codegen.PropertyType{
		Name:        "message",
		Description: lo.ToPtr("message at different levels, typically for error"),
	}
)

// app properties
var (
	PropertyTypeCardID = codegen.PropertyType{
		Name:        "card:id",
		Description: lo.ToPtr("card id"),
		Example:     lo.ToPtr("task:application:install"),
	}

	PropertyTypeCardBody = codegen.PropertyType{
		Name:        "card:body",
		Description: lo.ToPtr("card body"),
		Example:     lo.ToPtr("{xxxxxx}"),
	}
)

var (
	EventTypeYSKCardUpsert = codegen.EventType{
		SourceID: SERVICENAME,
		Name:     "ysk:card:upsert",
		PropertyTypeList: []codegen.PropertyType{
			PropertyTypeCardBody,
		},
	}

	EventTypeYSKCardDelete = codegen.EventType{
		SourceID: SERVICENAME,
		Name:     "ysk:card:delete",
		PropertyTypeList: []codegen.PropertyType{
			PropertyTypeCardID,
		},
	}
)
