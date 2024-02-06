package models

type Carriage struct {
	ID      uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    *string `json:"name"`
	Class   *string `json:"class"`
	TrainID *uint64 `json:"train_id"`
	Seats   []Seat  `gorm:"foreignKey:CarriageID" json:"seats"`
}
