package model

// TODO - add validation - see https://github.com/go-playground/validator

const PropertyTypeList = "PropertyTypeList"

type EventType struct {
	SourceID         string         `gorm:"primaryKey"`
	Name             string         `gorm:"primaryKey"`
	PropertyTypeList []PropertyType `gorm:"many2many:event_type_property_type;"`
}
