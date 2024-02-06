package models

type SeatPrice struct {
	ID         uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
    JourneyID  uint64  `json:"journey_id"`
    SeatID     uint64  `json:"seat_id"`
    Journey    Journey `json:"journey"`
    Seat       Seat    `json:"seat"`
    SeatPrice  uint64  `json:"seat_price"`
}