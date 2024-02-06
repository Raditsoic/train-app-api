package models

type BookedSeat struct {
	Seat    *Seat    `json:"seat"`
	Booking *Booking `json:"booking"`
}