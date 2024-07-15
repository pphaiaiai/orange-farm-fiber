package adapters

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pphaiaiai/orange-farm-fiber/entities"
	"github.com/pphaiaiai/orange-farm-fiber/usecases"
	"golang.org/x/crypto/bcrypt"
)

type HttpUserHandler struct {
	UserUseCase usecases.UserUseCase
}

func NewHttpUserHandler(useCase usecases.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{
		UserUseCase: useCase,
	}
}

func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	user.Password = string(hashedPassword)
	user.Role = "USER"

	err = h.UserUseCase.CreateUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "uni_users_email") {
			return c.Status(400).JSON(fiber.Map{"error": "Email already exists"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Registration successful"})
}
