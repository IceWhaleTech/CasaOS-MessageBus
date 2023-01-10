package service

import (
	"errors"

	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
)

type ActionTypeService struct {
	repository *repository.Repository
}

var (
	ErrActionSourceIDNotFound = errors.New("event source id not found")
	ErrActionNameNotFound     = errors.New("event name not found")
)

func (s *ActionTypeService) GetActionTypes() ([]model.ActionType, error) {
	return (*s.repository).GetActionTypes()
}

func (s *ActionTypeService) RegisterActionType(actionType model.ActionType) (*model.ActionType, error) {
	// TODO - ensure sourceID and name are URL safe

	return (*s.repository).RegisterActionType(actionType)
}

func (s *ActionTypeService) GetActionTypesBySourceID(sourceID string) ([]model.ActionType, error) {
	return (*s.repository).GetActionTypesBySourceID(sourceID)
}

func (s *ActionTypeService) GetActionType(sourceID string, name string) (*model.ActionType, error) {
	return (*s.repository).GetActionType(sourceID, name)
}

func NewActionTypeService(repository *repository.Repository) *ActionTypeService {
	return &ActionTypeService{
		repository: repository,
	}
}
