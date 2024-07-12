package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FarmRoute(app *fiber.App, db *gorm.DB) {

	app.Post("/farm", func(c *fiber.Ctx) error {
		return CreateFarm(db, c)
	})

	app.Get("/farm", func(c *fiber.Ctx) error {
		return GetFarms(db, c)
	})

	app.Get("/farm/:id", func(c *fiber.Ctx) error {
		return GetFarm(db, c)
	})

	app.Put("/farm/:id", func(c *fiber.Ctx) error {
		return UpdateFarm(db, c)
	})

	app.Delete("/farm/:id", func(c *fiber.Ctx) error {
		return DeleteFarm(db, c)
	})
}
