package service

import (
	"context"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
)

type YSKService struct{}

func NewYSKService() *YSKService {
	return &YSKService{}
}

func (s *YSKService) YskCardList(ctx *context.Context) (codegen.YSKCardList, error) {
	return nil, nil
}

func (s *YSKService) UpsertYSKCard(ctx *context.Context, yskCard codegen.YSKCard) error {
	return nil
}

func (s *YSKService) DeleteYSKCard(ctx *context.Context, id string) error {
	return nil
}
