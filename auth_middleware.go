package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func authRequired(c *fiber.Ctx) error {
	jwtSecretKey := os.Getenv("SECERT_KEY")
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)
	c.Locals("userID", int(userID))
	return c.Next()
}

func adminRequired(c *fiber.Ctx) error {
	jwtSecretKey := os.Getenv("SECERT_KEY")
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	log.Default().Print(role)
	if role != "ADMIN" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
	}
	userID := claims["user_id"].(float64)
	c.Locals("userID", int(userID))

	return c.Next()
}
