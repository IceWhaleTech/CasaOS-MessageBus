package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/common"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
	"go.uber.org/zap"
)

type EventService struct {
	ctx                *context.Context
	mutex              sync.Mutex
	repository         *repository.Repository
	inboundChannel     chan model.Event
	subscriberChannels map[string]map[string][]chan model.Event
	stop               chan struct{}
}

var (
	ErrEventSourceIDNotFound = errors.New("event source id not found")
	ErrEventNameNotFound     = errors.New("event name not found")
)

func (s *EventService) GetEventTypes() ([]model.EventType, error) {
	return (*s.repository).GetEventTypes()
}

func (s *EventService) RegisterEventType(eventType model.EventType) (*model.EventType, error) {
	// TODO - ensure sourceID and name are URL safe

	return (*s.repository).RegisterEventType(eventType)
}

func (s *EventService) GetEventTypesBySourceID(sourceID string) ([]model.EventType, error) {
	return (*s.repository).GetEventTypesBySourceID(sourceID)
}

func (s *EventService) GetEventType(sourceID string, name string) (*model.EventType, error) {
	return (*s.repository).GetEventType(sourceID, name)
}

func (s *EventService) Publish(event model.Event) (*model.Event, error) {
	if s.inboundChannel == nil {
		return nil, ErrInboundChannelNotFound
	}

	if event.Timestamp == 0 {
		event.Timestamp = time.Now().Unix()
	}

	// TODO - ensure properties are valid for event type

	select {
	case s.inboundChannel <- event:

	case <-(*s.ctx).Done():
		return nil, (*s.ctx).Err()

	default: // drop event if no one is listening
	}

	return &event, nil
}

func (s *EventService) Subscribe(sourceID string, names []string) (chan model.Event, error) {
	if len(names) == 0 {
		eventTypes, err := s.GetEventTypesBySourceID(sourceID)
		if err != nil {
			return nil, err
		}

		for _, eventType := range eventTypes {
			names = append(names, eventType.Name)
		}
	}

	for _, name := range names {
		eventType, err := s.GetEventType(sourceID, name)
		if err != nil {
			return nil, err
		}

		if eventType == nil {
			return nil, ErrEventNameNotFound
		}
	}

	if s.subscriberChannels == nil {
		s.subscriberChannels = make(map[string]map[string][]chan model.Event)
	}

	if s.subscriberChannels[sourceID] == nil {
		s.subscriberChannels[sourceID] = make(map[string][]chan model.Event)
	}

	c := make(chan model.Event, 1)

	for _, name := range names {
		if s.subscriberChannels[sourceID][name] == nil {
			s.subscriberChannels[sourceID][name] = make([]chan model.Event, 0)
		}
		s.subscriberChannels[sourceID][name] = append(s.subscriberChannels[sourceID][name], c)
	}

	return c, nil
}

func (s *EventService) Unsubscribe(sourceID string, name string, c chan model.Event) error {
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
		s.mutex.Lock()
		defer s.mutex.Unlock()

		if subscriber == c {
			logger.Info("unsubscribing from event type", zap.String("sourceID", sourceID), zap.String("name", name), zap.Int("subscriber", i))
			if i >= len(s.subscriberChannels[sourceID][name]) {
				logger.Error("the i-th subscriber is removed before we get here - concurrency issue?", zap.Int("subscriber", i), zap.Int("total", len(s.subscriberChannels[sourceID][name])))
				return ErrAlreadySubscribed
			}
			s.subscriberChannels[sourceID][name] = append(s.subscriberChannels[sourceID][name][:i], s.subscriberChannels[sourceID][name][i+1:]...)
			return nil
		}
	}

	return nil
}

func (s *EventService) Start(ctx *context.Context) {
	s.ctx = ctx
	s.mutex = sync.Mutex{}

	s.inboundChannel = make(chan model.Event)
	s.subscriberChannels = make(map[string]map[string][]chan model.Event)
	s.stop = make(chan struct{})

	defer func() {
		if s.subscriberChannels != nil {
			for sourceID, source := range s.subscriberChannels {
				for eventName, subscribers := range source {
					for _, subscriber := range subscribers {
						select {
						case _, ok := <-subscriber:
							if ok {
								close(subscriber)
							}
						default:
							continue
						}
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

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

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

		case <-ticker.C:
			if s.subscriberChannels == nil {
				continue
			}

			heartbeat := model.Event{
				SourceID:  common.MessageBusSourceID,
				Name:      common.MessageBusHeartbeatName,
				Timestamp: time.Now().Unix(),
			}

			for _, source := range s.subscriberChannels {
				for _, subscribers := range source {
					for _, subscriber := range subscribers {
						select {
						case subscriber <- heartbeat:
						case <-(*s.ctx).Done():
							return
						default: // drop event if no one is listening
							continue
						}
					}
				}
			}
		}
	}
}

func NewEventService(repository *repository.Repository) *EventService {
	return &EventService{
		repository: repository,
	}
}
