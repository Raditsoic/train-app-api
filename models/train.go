package models

type Train struct {
	ID   uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name *string `json:"name"`
}