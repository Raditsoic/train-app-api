package models

type SeatPrice struct {
	Journey   *Journey `json:"journey"`
	Seat      *Seat    `json:"seat"`
	SeatPrice *uint64  `json:"seat_price"`
}