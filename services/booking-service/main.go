package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/services/booking-service/database"
	"github.com/roshankumar18/event-booking/services/booking-service/kafka"
	"github.com/roshankumar18/event-booking/services/booking-service/routes"
)

func main() {
	err := kafka.InitKafkaProducer()
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %s", err)
	}
	database.InitDB()
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
