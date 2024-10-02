package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/roshankumar18/event-booking/services/user-service/database"
	"github.com/roshankumar18/event-booking/services/user-service/models"
	"github.com/roshankumar18/event-booking/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if c.ShouldBindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "name is required"})
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email is required"})
		return
	}
	if user.Role == models.Creator || user.Role == models.Booker {
		c.JSON(http.StatusBadRequest, gin.H{"message": "role is required"})
		return
	}
	var existingUser models.User
	userExists := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if userExists.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists"})
		return
	}
	hashPasswrod, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not hash password"})
		return
	}
	user.Password = string(hashPasswrod)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func LoginUser(c *gin.Context) {
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		var ve validator.ValidationErrors
		if ok := errors.As(err, &ve); ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "details": translateValidationErrors(ve)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}
	var existingUser models.User
	userExists := database.DB.Where("email = ?", inputUser.Email).First(&existingUser)
	if userExists.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(inputUser.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid password"})
		return
	}
	token, err := utils.GenerateToken(existingUser.ID, string(existingUser.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func translateValidationErrors(ve validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)
	for _, fe := range ve {
		errs[fe.Field()] = fe.Tag()
	}
	return errs
}
