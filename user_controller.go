package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoute(app *fiber.App, db *gorm.DB) {

	app.Post("/login", func(c *fiber.Ctx) error {
		return Login(db, c)
	})

	app.Get("/user/me", func(c *fiber.Ctx) error {
		return GetUserInfo(db, c)
	})
}
