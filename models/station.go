package models

type Station struct {
	ID   uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name *string `json:"name"`
	City *string `json:"city"`
}