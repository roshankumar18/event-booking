package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/roshankumar18/event-booking/pkg/database"
	"github.com/roshankumar18/event-booking/services/event-service/models"
	"github.com/roshankumar18/event-booking/utils"
)

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "details": utils.TranslateValidationErrors(ve)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}
	role, _ := c.Get("role")
	if role.(string) != "creator" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	userID, _ := c.Get("userID")
	event.CreatorID = userID.(uint)

	if err := database.DB.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "event created successfully",
		"event": event})

}

func GetEvent(c *gin.Context) {
	id := c.Params.ByName("id")
	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event"})
		return
	}
	//send event response to client
	c.JSON(http.StatusOK, gin.H{"event": event})
}

func UpdateEventSeats(eventID uint, seatsTaken int) error {
	var event models.Event
	if err := database.DB.First(&event, eventID).Error; err != nil {
		return err
	}
	event.SeatsAvailable -= seatsTaken
	if err := database.DB.Save(&event).Error; err != nil {
		return err
	}

	return nil
}

func GetAllEvent(c *gin.Context) {
	var events models.Event
	if err := database.DB.Where("seats_available > 0").Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})

}
