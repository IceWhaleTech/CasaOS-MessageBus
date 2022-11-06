package model

type Event struct {
	ID         uint       `gorm:"primaryKey"`
	SourceID   string     `gorm:"index"`
	Name       string     `gorm:"index"`
	Properties []Property `gorm:"foreignKey:Id"`
	Timestamp  int64      `gorm:"autoCreateTime:milli"`
}
