package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Variety struct {
	gorm.Model
	Name        string
	Description string
}

func GetVarieties(db *gorm.DB, c *fiber.Ctx) error {
	var varieties []Variety
	db.Find(&varieties)
	return c.JSON(varieties)
}

func GetVariety(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var variety Variety
	if db.First(&variety, id).Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Variety not found at ID: " + id})
	}
	return c.JSON(variety)
}

func CreateVariety(db *gorm.DB, c *fiber.Ctx) error {
	var variety Variety
	if err := c.BodyParser(&variety); err != nil {
		return err
	}
	db.Create(&variety)
	return c.JSON(variety)
}

func UpdateVariety(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	variety := new(Variety)
	db.First(&variety, id)
	if err := c.BodyParser(variety); err != nil {
		return err
	}
	db.Save(&variety)
	return c.JSON(variety)
}

func DeleteVariety(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	db.Delete(&Variety{}, id)
	return c.JSON(fiber.Map{"message": "Variety successfully deleted"})
}
