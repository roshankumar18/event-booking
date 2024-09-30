package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roshankumar18/event-booking/user-service/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/register", handlers.RegisterUser)
		user.POST("/login", handlers.LoginUser)
	}
}
