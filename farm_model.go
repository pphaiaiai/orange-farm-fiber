package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Farm struct {
	gorm.Model
	Name            string
	Location        string
	EstablishedDate string
	UserID          uint `json:"user_id" gorm:"unique"`
}

func CreateFarm(db *gorm.DB, c *fiber.Ctx) error {
	var farm Farm
	if err := c.BodyParser(&farm); err != nil {
		return err
	}

	userID, err := c.Locals("userID").(int)
	if !err {
		return c.Status(500).JSON(fiber.Map{"error": "User ID not found"})
	}
	farm.UserID = uint(userID)

	result := db.Create(&farm)
	if result.Error != nil {
		if strings.Contains((result.Error.Error()), "uni_farms_user_id") {
			return c.Status(500).JSON(fiber.Map{"error": "Farm already exists"})
		}
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(farm)
}

func GetFarms(db *gorm.DB, c *fiber.Ctx) error {
	var farms []Farm
	db.Find(&farms)
	return c.JSON(farms)
}

func GetFarm(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var farm Farm
	if db.First(&farm, id).Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Farm not found at ID: " + id})
	}
	return c.JSON(farm)
}

func UpdateFarm(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	farm := new(Farm)
	db.First(&farm, id)
	if err := c.BodyParser(farm); err != nil {
		return err
	}
	db.Save(&farm)
	return c.JSON(farm)
}

func DeleteFarm(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	db.Delete(&Farm{}, id)
	return c.JSON(fiber.Map{"message": "Farm successfully deleted"})
}
