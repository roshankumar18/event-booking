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
