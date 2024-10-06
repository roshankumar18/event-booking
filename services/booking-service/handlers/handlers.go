package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/roshankumar18/event-booking/services/booking-service/database"
	"github.com/roshankumar18/event-booking/services/booking-service/kafka"
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
	userEmail, _ := c.Get("email")
	booking := models.Booking{
		UserID:     userID.(uint),
		EventID:    bookingRequest.EventID,
		SeatsTaken: bookingRequest.SeatsTaken,
	}
	//check if seats are available
	seatsAvailable, err := checkSeatsAvailability(bookingRequest.EventID, bookingRequest.SeatsTaken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not check seat availability"})
		return
	}
	if !seatsAvailable {
		c.JSON(http.StatusBadRequest, gin.H{"message": "seats are not available"})
		return
	}

	if err := database.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create booking"})
		return
	}
	err = ProduceBookingMessage(userID.(uint), bookingRequest.EventID, bookingRequest.SeatsTaken, role.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not produce booking message"})
		return
	}
	go ProduceNotificationMessage(userID.(uint), bookingRequest.EventID, bookingRequest.SeatsTaken, userEmail.(string))
	c.JSON(http.StatusOK, gin.H{"message": "booking created successfully",
		"booking": booking})
}
func ProduceNotificationMessage(userId uint, eventId uint, seatsTaken int, email string) error {
	type NotificationMessage struct {
		UserID     uint   `json:"user_id"`
		EventID    uint   `json:"event_id"`
		SeatsTaken int    `json:"seats_taken"`
		Timestamp  int64  `json:"timestamp"`
		Email      string `json:"email"`
	}
	notificationMessage := NotificationMessage{
		UserID:     userId,
		EventID:    eventId,
		SeatsTaken: seatsTaken,
		Timestamp:  time.Now().Unix(),
		Email:      email,
	}
	json, err := json.Marshal(notificationMessage)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %v", err)
	}
	return kafka.ProduceMessage("booking-notifications", string(json))
}
func ProduceBookingMessage(userId, eventId uint, seatsTaken int, role string) error {
	type BookingMessage struct {
		UserID     uint   `json:"user_id"`
		EventID    uint   `json:"event_id"`
		SeatsTaken int    `json:"seats_taken"`
		Timestamp  int64  `json:"timestamp"`
		Role       string `json:"role"`
	}
	bookingMessage := BookingMessage{
		UserID:     userId,
		EventID:    eventId,
		SeatsTaken: seatsTaken,
		Timestamp:  time.Now().Unix(),
		Role:       role,
	}
	json, err := json.Marshal(bookingMessage)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %v", err)
	}
	return kafka.ProduceMessage("booking-service", string(json))
}

func checkSeatsAvailability(eventID uint, seatsNeeded int) (bool, error) {
	url := "http://" + os.Getenv("EVENT_SERVICE_URL") + "/events/?event_id=" + fmt.Sprint(eventID)
	return makeHTTPGetRequest(url, seatsNeeded)
}

func makeHTTPGetRequest(url string, seatsNeeded int) (bool, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var event struct {
		SeatsTaken int `json:"seats_taken"`
	}
	err = json.Unmarshal(body, &event)
	if err != nil {
		return false, err
	}

	return event.SeatsTaken >= seatsNeeded, nil
}
