package models

type Seat struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       *string   `json:"name"`
	Status     *string   `json:"status"`
	CarriageID uint64    `json:"carriage_id"`
	Carriage   Carriage  `json:"carriage"`
	SeatPrice  []Journey `gorm:"many2many:seat_prices;" json:"seat_price"`
	BookedSeat []Booking `gorm:"many2many:booked_ticket;" json:"book"`
}
