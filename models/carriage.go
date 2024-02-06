package models

type Carriage struct {
	ID    uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  *string `json:"name"`
	Class *string `json:"class"`
	Train *Train  `json:"train"`
}