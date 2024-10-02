package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/pkg/middleware"
	"github.com/roshankumar18/event-booking/services/booking-service/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/booking")
	{
		routes.Use(middleware.AuthMiddleware()).POST("/", handlers.CreateBooking)
		// user.POST("/login", handlers.LoginUser)
	}
}
