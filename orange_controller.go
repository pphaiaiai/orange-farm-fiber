package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func OrangeRoute(app *fiber.App, db *gorm.DB) {
	app.Get("/orange", func(c *fiber.Ctx) error {
		return GetOranges(db, c)
	})

	app.Get("/orange/:id", func(c *fiber.Ctx) error {
		return GetOrange(db, c)
	})

	app.Post("/orange", func(c *fiber.Ctx) error {
		return CreateOrange(db, c)
	})

	app.Put("/orange/:id", func(c *fiber.Ctx) error {
		return UpdateOrange(db, c)
	})

	app.Delete("/orange/:id", func(c *fiber.Ctx) error {
		return DeleteOrange(db, c)
	})
}
