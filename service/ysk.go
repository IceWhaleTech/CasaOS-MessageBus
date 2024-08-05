package service

import (
	"context"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type YSKService struct {
	repository *repository.Repository
}

func NewYSKService(repository *repository.Repository) *YSKService {
	return &YSKService{
		repository: repository,
	}
}

func (s *YSKService) YskCardList(ctx context.Context) (codegen.YSKCardList, error) {
	cardList, err := (*s.repository).GetYSKCardList()
	if err != nil {
		return codegen.YSKCardList{}, err
	}
	return cardList, nil
}

func (s *YSKService) UpsertYSKCard(ctx context.Context, yskCard codegen.YSKCard) error {
	err := (*s.repository).UpsertYSKCard(yskCard)
	return err
}

func (s *YSKService) DeleteYSKCard(ctx context.Context, id string) error {
	return nil
}
