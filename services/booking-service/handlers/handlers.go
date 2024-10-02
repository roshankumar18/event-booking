package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/roshankumar18/event-booking/services/booking-service/database"
	"github.com/roshankumar18/event-booking/services/booking-service/models"
	"github.com/roshankumar18/event-booking/utils"
)

type BookingRequest struct {
	EventID    uint `json:"event_id" binding:"required"`
	SeatsTaken int  `json:"seats_taken" binding:"required"`
}

func CreateBooking(c *gin.Context) {
	var bookingRequest BookingRequest
	if err := c.ShouldBindJSON(&bookingRequest); err != nil {
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "details": utils.TranslateValidationErrors(ve)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}
	role, _ := c.Get("role")
	if role.(string) != "booker" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	userID, _ := c.Get("userID")
	booking := models.Booking{
		UserID:     userID.(uint),
		EventID:    bookingRequest.EventID,
		SeatsTaken: bookingRequest.SeatsTaken,
	}

	if err := database.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create booking"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "booking created successfully",
		"booking": booking})
}
