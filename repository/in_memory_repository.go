package repository

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type InMemoryRepository struct {
	db *gorm.DB
}

func (r *InMemoryRepository) GetEventTypes() ([]model.EventType, error) {
	var eventTypes []model.EventType

	if err := r.db.Preload(model.PropertyTypeList).Find(&eventTypes).Error; err != nil {
		return nil, err
	}

	return eventTypes, nil
}

func (r *InMemoryRepository) RegisterEventType(eventType model.EventType) (*model.EventType, error) {
	if err := r.db.Create(&eventType).Error; err != nil {
		return nil, err
	}

	return &eventType, nil
}

func (r *InMemoryRepository) GetEventTypesBySourceID(sourceID string) ([]model.EventType, error) {
	var eventTypes []model.EventType

	if err := r.db.Preload(model.PropertyTypeList).Where(&model.EventType{SourceID: sourceID}).Find(&eventTypes).Error; err != nil {
		return nil, err
	}

	return eventTypes, nil
}

func (r *InMemoryRepository) GetEventType(sourceID string, name string) (*model.EventType, error) {
	var eventType model.EventType

	if err := r.db.Preload(model.PropertyTypeList).Where(&model.EventType{SourceID: sourceID, Name: name}).First(&eventType).Error; err != nil {
		return nil, err
	}

	return &eventType, nil
}

func (r *InMemoryRepository) Close() {
	sqlDB, err := r.db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func NewInMemoryRepository() (Repository, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.EventType{}, &model.PropertyType{}); err != nil {
		return nil, err
	}

	return &InMemoryRepository{
		db: db,
	}, nil
}
