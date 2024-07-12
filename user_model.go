package main

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     Role   `json:"role"`
}

func CreateUser(db *gorm.DB, c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	user.Password = string(hashedPassword)
	user.Role = "USER"

	result := db.Create(user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "uni_users_email") {
			return c.Status(500).JSON(fiber.Map{"error": "Email already exists"})
		}
		return c.Status(500).SendString(result.Error.Error())
	}

	return c.JSON(fiber.Map{"message": "Registration successful"})
}

// loginUser handles user login
func Login(db *gorm.DB, c *fiber.Ctx) error {
	var input User
	var user User

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	// Find user by email
	db.Where("email = ?", input.Email).First(&user)

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = user.Role

	jwtSecretKey := os.Getenv("SECERT_KEY")
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "Login successful"})
}

func GetUserInfo(db *gorm.DB, c *fiber.Ctx) error {
	userID, err := c.Locals("userID").(int)
	if !err {
		return c.Status(500).JSON(fiber.Map{"error": "User ID not found"})
	}

	var user User
	db.First(&user, userID)
	user.Password = ""

	return c.JSON(user)
}
