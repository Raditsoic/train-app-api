package models

import "time"

type Journey struct {
	ID               uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             *string    `json:"name"`
	ArrivalTime      *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"arrival_time"`
	DepartureTime    *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"departure_time"`
	Train            *Train     `json:"train"`
	ArrivalStation   *Station   `json:"arrival_station"`
	DepartureStation *Station   `json:"departure_station"`
}