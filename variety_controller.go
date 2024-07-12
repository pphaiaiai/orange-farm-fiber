package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func VarietyRoute(app *fiber.App, db *gorm.DB) {
	app.Get("/variety", func(c *fiber.Ctx) error {
		return (GetVarieties(db, c))
	})

	app.Get("/variety/:id", func(c *fiber.Ctx) error {
		return GetVariety(db, c)
	})

	app.Post("/variety", func(c *fiber.Ctx) error {
		return CreateVariety(db, c)
	})

	app.Put("/variety/:id", func(c *fiber.Ctx) error {
		return UpdateVariety(db, c)
	})

	app.Delete("/variety/:id", func(c *fiber.Ctx) error {
		return DeleteVariety(db, c)
	})
}
