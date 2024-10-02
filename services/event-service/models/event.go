package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title          string    `json:"title" binding:"required"`
	Description    string    `json:"description"`
	StartTime      time.Time `json:"start_time" binding:"required"`
	EndTime        time.Time `json:"end_time" binding:"required"`
	Location       string    `json:"location" binding:"required"`
	CreatorID      uint      `json:"creator_id"`
	SeatsAvailable int       `json:"seats_available" binding:"required"`
	Price          int       `json:"price"`
}
