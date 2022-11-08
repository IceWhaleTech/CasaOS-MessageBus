package repository

import (
	"time"

	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseRepository struct {
	db *gorm.DB
}

func (r *DatabaseRepository) GetEventTypes() ([]model.EventType, error) {
	var eventTypes []model.EventType

	if err := r.db.Preload(model.PropertyTypeList).Find(&eventTypes).Error; err != nil {
		return nil, err
	}

	return eventTypes, nil
}

func (r *DatabaseRepository) RegisterEventType(eventType model.EventType) (*model.EventType, error) {
	if err := r.db.Create(&eventType).Error; err != nil {
		return nil, err
	}

	return &eventType, nil
}

func (r *DatabaseRepository) GetEventTypesBySourceID(sourceID string) ([]model.EventType, error) {
	var eventTypes []model.EventType

	if err := r.db.Preload(model.PropertyTypeList).Where(&model.EventType{SourceID: sourceID}).Find(&eventTypes).Error; err != nil {
		return nil, err
	}

	return eventTypes, nil
}

func (r *DatabaseRepository) GetEventType(sourceID string, name string) (*model.EventType, error) {
	var eventType model.EventType

	if err := r.db.Preload(model.PropertyTypeList).Where(&model.EventType{SourceID: sourceID, Name: name}).First(&eventType).Error; err != nil {
		return nil, err
	}

	return &eventType, nil
}

func (r *DatabaseRepository) Close() {
	sqlDB, err := r.db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func NewDatabaseRepositoryInMemory() (Repository, error) {
	return NewDatabaseRepository("file::memory:?cache=shared")
}

func NewDatabaseRepository(databaseFilePath string) (Repository, error) {
	db, err := gorm.Open(sqlite.Open(databaseFilePath))
	if err != nil {
		return nil, err
	}

	c, _ := db.DB()
	c.SetMaxIdleConns(10)
	c.SetMaxOpenConns(100)
	c.SetConnMaxIdleTime(1000 * time.Second)

	if err := db.AutoMigrate(&model.EventType{}, &model.PropertyType{}); err != nil {
		return nil, err
	}

	return &DatabaseRepository{
		db: db,
	}, nil
}
