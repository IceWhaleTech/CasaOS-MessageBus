package service

import (
	"errors"

	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type EventTypeService struct {
	repository *repository.Repository
}

var (
	ErrEventSourceIDNotFound = errors.New("event source id not found")
	ErrEventNameNotFound     = errors.New("event name not found")
)

func (s *EventTypeService) GetEventTypes() ([]model.EventType, error) {
	return (*s.repository).GetEventTypes()
}

func (s *EventTypeService) RegisterEventType(eventType model.EventType) (*model.EventType, error) {
	// TODO - ensure sourceID and name are URL safe

	return (*s.repository).RegisterEventType(eventType)
}

func (s *EventTypeService) GetEventTypesBySourceID(sourceID string) ([]model.EventType, error) {
	return (*s.repository).GetEventTypesBySourceID(sourceID)
}

func (s *EventTypeService) GetEventType(sourceID string, name string) (*model.EventType, error) {
	return (*s.repository).GetEventType(sourceID, name)
}

func NewEventTypeService(repository *repository.Repository) *EventTypeService {
	return &EventTypeService{
		repository: repository,
	}
}
