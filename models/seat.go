package models

type Seat struct {
	ID       uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     *string   `json:"name"`
	Status   *string   `json:"status"`
	Carriage *Carriage `json:"carriage"`
}