package service

import (
	"context"
	"errors"
	"time"

	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type EventTypeService struct {
	ctx                *context.Context
	repository         repository.Repository
	inboundChannel     chan model.Event
	subscriberChannels map[string]map[string][]chan model.Event
	stop               chan struct{}
}

var (
	ErrInboundChannelNotFound     = errors.New("inbound channel not found")
	ErrSubscriberChannelsNotFound = errors.New("subscriber channels not found")
	ErrEventSourceIDNotFound      = errors.New("event source id not found")
	ErrEventNameNotFound          = errors.New("event name not found")
)

func (s *EventTypeService) GetEventTypes() ([]model.EventType, error) {
	return s.repository.GetEventTypes()
}

func (s *EventTypeService) RegisterEventType(eventType model.EventType) (*model.EventType, error) {
	return s.repository.RegisterEventType(eventType)
}

func (s *EventTypeService) GetEventTypesBySourceID(sourceID string) ([]model.EventType, error) {
	return s.repository.GetEventTypesBySourceID(sourceID)
}

func (s *EventTypeService) GetEventType(sourceID string, name string) (*model.EventType, error) {
	return s.repository.GetEventType(sourceID, name)
}

func (s *EventTypeService) Publish(event model.Event) (*model.Event, error) {
	if s.inboundChannel == nil {
		return nil, ErrInboundChannelNotFound
	}

	if event.Timestamp == 0 {
		event.Timestamp = time.Now().Unix()
	}

	select {
	case s.inboundChannel <- event:

	case <-(*s.ctx).Done():
		return nil, (*s.ctx).Err()

	default: // drop event if no one is listening
	}

	return &event, nil
}

func (s *EventTypeService) Subscribe(sourceID string, name string) (chan model.Event, error) {
	eventType, err := s.GetEventType(sourceID, name)
	if err != nil {
		return nil, err
	}

	if eventType == nil {
		return nil, ErrEventNameNotFound
	}

	if s.subscriberChannels == nil {
		s.subscriberChannels = make(map[string]map[string][]chan model.Event)
	}

	if s.subscriberChannels[sourceID] == nil {
		s.subscriberChannels[sourceID] = make(map[string][]chan model.Event)
	}

	if s.subscriberChannels[sourceID][name] == nil {
		s.subscriberChannels[sourceID][name] = make([]chan model.Event, 0)
	}

	c := make(chan model.Event, 1)
	s.subscriberChannels[sourceID][name] = append(s.subscriberChannels[sourceID][name], c)

	return c, nil
}

func (s *EventTypeService) Unsubscribe(sourceID string, name string, c chan model.Event) error {
	if s.subscriberChannels == nil {
		return ErrSubscriberChannelsNotFound
	}

	if s.subscriberChannels[sourceID] == nil {
		return ErrEventSourceIDNotFound
	}

	if s.subscriberChannels[sourceID][name] == nil {
		return ErrEventNameNotFound
	}

	for i, subscriber := range s.subscriberChannels[sourceID][name] {
		if subscriber == c {
			s.subscriberChannels[sourceID][name] = append(s.subscriberChannels[sourceID][name][:i], s.subscriberChannels[sourceID][name][i+1:]...)
			close(c)
			return nil
		}
	}

	return nil
}

func (s *EventTypeService) Start(ctx *context.Context) {
	s.ctx = ctx

	s.inboundChannel = make(chan model.Event)
	s.subscriberChannels = make(map[string]map[string][]chan model.Event)
	s.stop = make(chan struct{})

	defer func() {
		if s.subscriberChannels != nil {
			for sourceID, source := range s.subscriberChannels {
				for eventName, subscribers := range source {
					for _, subscriber := range subscribers {
						close(subscriber)
					}
					delete(s.subscriberChannels[sourceID], eventName)
				}
				delete(s.subscriberChannels, sourceID)
			}
			s.subscriberChannels = nil
		}

		close(s.inboundChannel)
		s.inboundChannel = nil

		close(s.stop)
		s.stop = nil
	}()

	for {
		select {

		case <-(*s.ctx).Done():
			return

		case event, ok := <-s.inboundChannel:
			if !ok {
				return
			}

			if s.subscriberChannels == nil {
				continue
			}

			if s.subscriberChannels[event.SourceID] == nil {
				continue
			}

			if s.subscriberChannels[event.SourceID][event.Name] == nil {
				continue
			}

			for _, c := range s.subscriberChannels[event.SourceID][event.Name] {
				select {
				case c <- event:
				case <-(*s.ctx).Done():
					return
				default: // drop event if no one is listening
					continue
				}
			}
		}
	}
}

func NewEventTypeService(repository repository.Repository) EventTypeService {
	return EventTypeService{
		repository: repository,
	}
}
