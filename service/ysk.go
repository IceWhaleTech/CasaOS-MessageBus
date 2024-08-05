package service

import (
	"context"
	"fmt"
	"log"

	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type YSKService struct {
	repository *repository.Repository
	ws         *EventServiceWS
}

func NewYSKService(repository *repository.Repository, ws *EventServiceWS) *YSKService {
	return &YSKService{
		repository: repository,
		ws:         ws,
	}
}

func (s *YSKService) YskCardList(ctx context.Context) ([]ysk.YSKCard, error) {
	cardList, err := (*s.repository).GetYSKCardList()
	if err != nil {
		return []ysk.YSKCard{}, err
	}
	return cardList, nil
}

func (s *YSKService) UpsertYSKCard(ctx context.Context, yskCard ysk.YSKCard) error {
	// don't store short notice cards
	if yskCard.CardType == ysk.CardTypeShortNote {
		return nil
	}
	err := (*s.repository).UpsertYSKCard(yskCard)
	return err
}

func (s *YSKService) DeleteYSKCard(ctx context.Context, id string) error {
	return nil
}

func (s *YSKService) Start() {
	channel, err := s.ws.Subscribe(ysk.SERVICENAME, []string{
		"ysk:card:create", "ysk:card:delete",
	})
	if err != nil {
		return
	}

	go func() {
		for {
			select {
			case event, ok := <-channel:
				if !ok {
					log.Println("channel closed")
				}
				fmt.Println(event)
			}
		}
	}()

}
