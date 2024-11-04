package model

type Settings struct {
	Key   string `gorm:"primaryKey"`
	Value string
}
