package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/pkg/database"
	"github.com/roshankumar18/event-booking/pkg/kafka"
	"github.com/roshankumar18/event-booking/services/booking-service/models"
	"github.com/roshankumar18/event-booking/services/booking-service/routes"
	"github.com/roshankumar18/event-booking/utils"
)

func main() {
	err := kafka.InitKafkaProducer()
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %s", err)
	}
	database.InitDB(utils.GoDotEnvVariable("DB_CONNECTION_STRING"), &models.Booking{})
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
