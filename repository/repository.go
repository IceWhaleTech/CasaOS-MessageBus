package repository

import "github.com/IceWhaleTech/CasaOS-MessageBus/model"

type Repository interface {
	GetEventTypes() ([]model.EventType, error)
	RegisterEventType(eventType model.EventType) (*model.EventType, error)
	GetEventTypesBySourceID(sourceID string) ([]model.EventType, error)
	GetEventType(sourceID string, name string) (*model.EventType, error)
	Close()
}
