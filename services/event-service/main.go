package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/pkg/database"
	"github.com/roshankumar18/event-booking/services/event-service/handlers"
	"github.com/roshankumar18/event-booking/services/event-service/models"
	"github.com/roshankumar18/event-booking/services/event-service/routes"
	"github.com/roshankumar18/event-booking/utils"
)

type BookingMessage struct {
	UserID     uint   `json:"user_id"`
	EventID    uint   `json:"event_id"`
	SeatsTaken int    `json:"seats_taken"`
	Timestamp  int64  `json:"timestamp"`
	Role       string `json:"role"`
}

func main() {
	database.InitDB(utils.GoDotEnvVariable("DB_CONNECTION_STRING"), &models.Event{})
	router := gin.Default()
	routes.RegisterRoutes(router)
	go startConsumer()
	router.Run()
}

func startConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "event-service-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	c.Subscribe("booking-service", nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var updateMessage BookingMessage
			if err := json.Unmarshal(msg.Value, &updateMessage); err != nil {
				fmt.Printf("Error decoding message: %v\n", err)
				continue
			}

			err := handlers.UpdateEventSeats(updateMessage.EventID, updateMessage.SeatsTaken)
			if err != nil {
				fmt.Printf("Error updating event seats: %v\n", err)
			} else {
				fmt.Printf("Event seats updated successfully\n")
			}
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
