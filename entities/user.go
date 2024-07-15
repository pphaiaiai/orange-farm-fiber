package entities

import "gorm.io/gorm"

type Role string

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     Role   `json:"role"`
}
