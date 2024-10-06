package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/pkg/database"
	"github.com/roshankumar18/event-booking/services/event-service/models"
	"github.com/roshankumar18/event-booking/services/event-service/routes"
	"github.com/roshankumar18/event-booking/utils"
)

func main() {
	database.InitDB(utils.GoDotEnvVariable("DB_CONNECTION_STRING"), &models.Event{})
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
