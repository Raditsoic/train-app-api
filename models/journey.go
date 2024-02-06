package models

import "time"

type Journey struct {
	ID                 uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name               *string     `json:"name"`
	ArrivalTime        *time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"arrival_time"`
	DepartureTime      *time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"departure_time"`
	TrainID            uint64      `json:"train_id"`
	Train              Train       `json:"train"`
	ArrivalStationID   uint64      `json:"arrival_station_id"`
	ArrivalStation     Station     `json:"arrival_station"`
	DepartureStationID uint64      `json:"departure_station_id"`
	DepartureStation   Station     `json:"departure_station"`
	Stations           []Station   `gorm:"many2many:journey_stations;" json:"stations"`
	SeatPrices         []SeatPrice `json:"seat_prices"`
}
