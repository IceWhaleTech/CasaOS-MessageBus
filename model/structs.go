package model

const (
	EventTypeList    = "EventTypeList"
	ActionTypeList   = "ActionTypeList"
	PropertyTypeList = "PropertyTypeList"
)

type EventType struct {
	SourceID         string         `gorm:"primaryKey"`
	Name             string         `gorm:"primaryKey"`
	PropertyTypeList []PropertyType `gorm:"many2many:event_type_property_type;"`
}

type Event struct {
	ID         uint              `gorm:"primaryKey"`
	SourceID   string            `gorm:"index"`
	Name       string            `gorm:"index"`
	Properties map[string]string `gorm:"foreignKey:Id"`
	Timestamp  int64             `gorm:"autoCreateTime:milli"`
	UUID       string            `json:"uuid,omitempty"`
}

type ActionType struct {
	SourceID         string         `gorm:"primaryKey"`
	Name             string         `gorm:"primaryKey"`
	PropertyTypeList []PropertyType `gorm:"many2many:action_type_property_type;"`
}

type Action struct {
	ID         uint              `gorm:"primaryKey"`
	SourceID   string            `gorm:"index"`
	Name       string            `gorm:"index"`
	Properties map[string]string `gorm:"foreignKey:Id"`
	Timestamp  int64             `gorm:"autoCreateTime:milli"`
}

type PropertyType struct {
	Name string `gorm:"primaryKey"`
}

// type Property struct {
// 	ID    uint `gorm:"primaryKey"`
// 	Name  string
// 	Value string
// }

type GenericType struct {
	SourceID string `gorm:"primaryKey"`
	Name     string `gorm:"primaryKey"`
}
