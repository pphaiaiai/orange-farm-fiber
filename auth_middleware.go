package main

import (
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
