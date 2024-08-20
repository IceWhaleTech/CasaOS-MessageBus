package repository

import (
	"os"
	"path/filepath"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DatabaseRepository struct {
	db *gorm.DB
}

func (r *DatabaseRepository) GetEvent(id string) (*model.Event, error) {
	var event model.Event
	if err := r.db.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *DatabaseRepository) GetEvents(sourceID string, eventType string) ([]model.Event, error) {
	var events []model.Event
	query := r.db.Where("source_id = ?", sourceID)
	if eventType != "" {
		query = query.Where("event_type = ?", eventType)
	}
	if err := query.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *DatabaseRepository) InsertEvent(event model.Event) error {
	return r.db.Create(&event).Error
}

func (r *DatabaseRepository) GetYSKCardList() ([]ysk.YSKCard, error) {
	var cardList []ysk.YSKCard
	if err := r.db.Find(&cardList).Error; err != nil {
		return nil, err
	}
	return cardList, nil
}

func (r *DatabaseRepository) UpsertYSKCard(card ysk.YSKCard) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&card).Error
}
func (r *DatabaseRepository) DeleteYSKCard(id string) error {
	return r.db.Delete(&ysk.YSKCard{Id: id}).Error
}
func (r *DatabaseRepository) GetEventTypes() ([]model.EventType, error) {
	var eventTypes []model.EventType

	if err := r.db.Preload(model.PropertyTypeList).Find(&eventTypes).Error; err != nil {
		return nil, err
	}

	return eventTypes, nil
}

func (r *DatabaseRepository) RegisterEventType(eventType model.EventType) (*model.EventType, error) {
	// upsert
	if err := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&eventType).Error; err != nil {
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
		logger.Error("can't find event type", zap.String("sourceID", sourceID), zap.String("EventName", name), zap.Error(err))
		return nil, err
	}

	return &eventType, nil
}

func (r *DatabaseRepository) GetActionTypes() ([]model.ActionType, error) {
	return GetTypes[model.ActionType](r.db)
}

func (r *DatabaseRepository) RegisterActionType(actionType model.ActionType) (*model.ActionType, error) {
	return RegisterType(r.db, actionType)
}

func (r *DatabaseRepository) GetActionTypesBySourceID(sourceID string) ([]model.ActionType, error) {
	return GetTypesBySourceID[model.ActionType](r.db, sourceID)
}

func (r *DatabaseRepository) GetActionType(sourceID string, name string) (*model.ActionType, error) {
	return GetType[model.ActionType](r.db, sourceID, name)
}

func (r *DatabaseRepository) Close() {
	sqlDB, err := r.db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func GetTypes[T any](db *gorm.DB) ([]T, error) {
	var types []T

	if err := db.Preload(model.PropertyTypeList).Find(&types).Error; err != nil {
		return nil, err
	}

	return types, nil
}

func RegisterType[T any](db *gorm.DB, t T) (*T, error) {
	// upsert
	if err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&t).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

func GetTypesBySourceID[T any](db *gorm.DB, sourceID string) ([]T, error) {
	var types []T

	if err := db.Preload(model.PropertyTypeList).Where(&model.GenericType{SourceID: sourceID}).Find(&types).Error; err != nil {
		return nil, err
	}

	return types, nil
}

func GetType[T any](db *gorm.DB, sourceID string, name string) (*T, error) {
	var t T

	if err := db.Preload(model.PropertyTypeList).Where(&model.GenericType{SourceID: sourceID, Name: name}).First(&t).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

func NewDatabaseRepositoryInMemory() (Repository, error) {
	return NewDatabaseRepository("file::memory:?cache=shared")
}

func NewDatabaseRepository(databaseFilePath string) (Repository, error) {
	// mkdir dbpath, 777 is copied from zimaos-local-storage
	if err := os.MkdirAll(filepath.Dir(databaseFilePath), 0o777); err != nil {
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(databaseFilePath))
	if err != nil {
		return nil, err
	}

	c, _ := db.DB()
	c.SetMaxIdleConns(10)
	c.SetMaxOpenConns(1)
	c.SetConnMaxIdleTime(1000 * time.Second)

	if err := db.AutoMigrate(
		&model.EventType{}, &model.ActionType{}, &model.PropertyType{},
		&ysk.YSKCard{},
	); err != nil {
		return nil, err
	}

	return &DatabaseRepository{
		db: db,
	}, nil
}
