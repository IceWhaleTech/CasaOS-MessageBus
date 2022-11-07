package repository

import "gorm.io/gorm"

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(databasePath string) *DBRepository {
	return &DBRepository{db: db}
}
