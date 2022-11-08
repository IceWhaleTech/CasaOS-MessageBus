package model

type Property struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Value string
}
