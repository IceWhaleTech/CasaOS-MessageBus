package service

import (
	"context"
	"sync"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/common"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"go.uber.org/zap"
)

type ActionServiceWS struct {
	typeService *ActionTypeService

	ctx   *context.Context
	mutex sync.Mutex
	stop  chan struct{}

	inboundChannel     chan model.Action
	subscriberChannels map[string]map[string][]chan model.Action
}

func (s *ActionServiceWS) Trigger(action model.Action) {
	if s.inboundChannel == nil {
		logger.Error("error when triggering action via websocket", zap.Error(ErrInboundChannelNotFound))
	}

	if action.Timestamp == 0 {
		action.Timestamp = time.Now().Unix()
	}

	// TODO - ensure properties are valid for action type

	select {
	case s.inboundChannel <- action:

	case <-(*s.ctx).Done():
		if err := (*s.ctx).Err(); err != nil {
			logger.Info(err.Error())
		}
		return

	default: // drop action if no one is listening
	}
}

func (s *ActionServiceWS) Subscribe(sourceID string, names []string) (chan model.Action, error) {
	if len(names) == 0 {
		actionTypes, err := s.typeService.GetActionTypesBySourceID(sourceID)
		if err != nil {
			return nil, err
		}

		for _, actionType := range actionTypes {
			names = append(names, actionType.Name)
		}
	}

	for _, name := range names {
		actionType, err := s.typeService.GetActionType(sourceID, name)
		if err != nil {
			return nil, err
		}

		if actionType == nil {
			return nil, ErrActionNameNotFound
		}
	}

	if s.subscriberChannels == nil {
		s.subscriberChannels = make(map[string]map[string][]chan model.Action)
	}

	if s.subscriberChannels[sourceID] == nil {
		s.subscriberChannels[sourceID] = make(map[string][]chan model.Action)
	}

	c := make(chan model.Action, 1)

	for _, name := range names {
		if s.subscriberChannels[sourceID][name] == nil {
			s.subscriberChannels[sourceID][name] = make([]chan model.Action, 0)
		}
		s.subscriberChannels[sourceID][name] = append(s.subscriberChannels[sourceID][name], c)
	}

	return c, nil
}

func (s *ActionServiceWS) Unsubscribe(sourceID string, name string, c chan model.Action) error {
	if s.subscriberChannels == nil {
		return ErrSubscriberChannelsNotFound
	}

	if s.subscriberChannels[sourceID] == nil {
		return ErrActionSourceIDNotFound
	}

	if s.subscriberChannels[sourceID][name] == nil {
		return ErrActionNameNotFound
	}

	for i, subscriber := range s.subscriberChannels[sourceID][name] {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		if subscriber == c {
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

func (s *ActionServiceWS) Start(ctx *context.Context) {
	s.ctx = ctx
	s.mutex = sync.Mutex{}

	s.inboundChannel = make(chan model.Action)
	s.subscriberChannels = make(map[string]map[string][]chan model.Action)
	s.stop = make(chan struct{})

	defer func() {
		if s.subscriberChannels != nil {
			for sourceID, source := range s.subscriberChannels {
				for actionName, subscribers := range source {
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
					delete(s.subscriberChannels[sourceID], actionName)
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

		case action, ok := <-s.inboundChannel:
			if !ok {
				return
			}

			if s.subscriberChannels == nil {
				continue
			}

			if s.subscriberChannels[action.SourceID] == nil {
				continue
			}

			if s.subscriberChannels[action.SourceID][action.Name] == nil {
				continue
			}

			for _, c := range s.subscriberChannels[action.SourceID][action.Name] {
				select {
				case c <- action:
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

			heartbeat := model.Action{
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

func NewActionServiceWS(actionTypeService *ActionTypeService) *ActionServiceWS {
	return &ActionServiceWS{
		typeService: actionTypeService,
	}
}
