package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/services/booking-service/database"
	"github.com/roshankumar18/event-booking/services/booking-service/routes"
)

func main() {
	database.InitDB()
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
