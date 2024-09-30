package models

import "gorm.io/gorm"

type Role string

const (
	Creator Role = "creator"
	Booker  Role = "booker"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email" binding:"required,email"`
	Password string `json:"password"`
	Role     Role   `json:"role" binding:"required"`
}
