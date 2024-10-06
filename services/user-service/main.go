package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/pkg/database"
	"github.com/roshankumar18/event-booking/services/user-service/models"
	"github.com/roshankumar18/event-booking/services/user-service/routes"
	"github.com/roshankumar18/event-booking/utils"
)

func main() {
	database.InitDB(utils.GoDotEnvVariable("DB_CONNECTION_STRING"), &models.User{})
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
