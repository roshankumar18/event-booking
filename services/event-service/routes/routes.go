package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/pkg/middleware"
	"github.com/roshankumar18/event-booking/services/event-service/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	user := router.Group("/events")
	{
		user.Use(middleware.AuthMiddleware()).POST("/", handlers.CreateEvent).GET("/:id", handlers.GetEvent)
	}
}
