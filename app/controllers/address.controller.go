package controllers

import (
	"app-transaction/database"
	"app-transaction/models"

	"github.com/gofiber/fiber/v2"
)

func AddressCreate(c *fiber.Ctx) error {

	address := new(models.Address)

	if err := c.BodyParser(address); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	database.DB.Create(&address)

	return c.JSON(fiber.Map{
		"message": "create customer successfully",
		"post":    address,
	})

}

func AddressGetAll(c *fiber.Ctx) error {
	var address []models.Address

	err := database.DB.Preload("Address").Find(&address)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return c.JSON(fiber.Map{
		"customer": address,
	})
}

func AddressGetByID(c *fiber.Ctx) error {
	addressID := c.Params("ID")

	var address []models.Address

	err := database.DB.First(&address, "ID = ? ", addressID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    address,
	})
}

// func CustomerUpdate(c *fiber.Ctx) error {

// }

func AddressDelete(c *fiber.Ctx) error {
	addressID := c.Params("ID")

	var address []models.Customer

	err := database.DB.Debug().First(&address, "ID = ?", addressID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&address).Error

	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "address was deleted",
	})

}
