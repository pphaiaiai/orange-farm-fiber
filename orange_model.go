package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Orange struct {
	gorm.Model
	VarietyId   uint
	PlantDate   string
	HarvestDate string
	Quantity    int
	FarmID      uint `json:"farm_id"`
}

func GetOranges(db *gorm.DB, c *fiber.Ctx) error {
	var oranges []Orange

	userID, err := c.Locals("userID").(int)
	if !err {
		return c.Status(500).JSON(fiber.Map{"error": "User ID not found"})
	}

	var farm Farm
	db.Where("user_id = ?", userID).First(&farm)
	db.Find(&oranges, "farm_id = ?", farm.ID)
	return c.JSON(oranges)
}

func GetOrange(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var orange Orange
	db.Find(&orange, id)
	return c.JSON(orange)
}

func CreateOrange(db *gorm.DB, c *fiber.Ctx) error {
	var orange Orange

	if err := c.BodyParser(&orange); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	userID, err := c.Locals("userID").(int)
	if !err {
		return c.Status(500).JSON(fiber.Map{"error": "User ID not found"})
	}

	var farm Farm
	db.Where("user_id = ?", userID).First(&farm)
	orange.FarmID = farm.ID

	db.Create(&orange)
	return c.JSON(orange)
}

func UpdateOrange(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	orange := new(Orange)
	if err := c.BodyParser(orange); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Find(&orange, id)
	db.Save(&orange)
	return c.JSON(orange)
}

func DeleteOrange(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var orange Orange
	db.Find(&orange, id)
	db.Delete(&orange)
	return c.SendString("Orange successfully deleted")
}
