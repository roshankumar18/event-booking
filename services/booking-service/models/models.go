package models

import "gorm.io/gorm"

type Status string

const (
	Success Status = "success"
	Failed  Status = "failed"
)

type Booking struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	EventID    uint    `json:"event_id"`
	Status     Status  `json:"status"`
	SeatsTaken int     `json:"seats_taken"`
	Price      float64 `json:"price"`
}
