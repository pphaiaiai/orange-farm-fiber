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

// Handler functions
// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with email and password
// @Tags user
// @Accept  json
// @Produce  json
// @Security
// @Success 200 {array} User
// @Router /register [post]
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

// Login godoc
// @Summary Login for authentication
// @Description Login with email and password to get JWT token
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body User true "User credentials" example({"email": "user@example.com", "password": "test"})
// @Success 200 {object} map[string]string "Login successful"
// @Failure 401 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /login [post]
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
