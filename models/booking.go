package models

import "time"

type Booking struct {
	ID            uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	BookDate      *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"book_date"`
	PaymentStatus bool       `gorm:"default:false" json:"payment_status"`
	PaymentType   string     `json:"payment_type"`
	PaymentDate   time.Time  `json:"payment_date"`
	Seats         []Seat     `gorm:"many2many:booked_seats" json:"booked_seats"`
}
