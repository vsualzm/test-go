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
	var customer []models.Customer

	err := database.DB.Preload("Customer").Find(&customer)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return c.JSON(fiber.Map{
		"customer": customer,
	})
}

func AddressGetByID(c *fiber.Ctx) error {
	customerID := c.Params("ID")

	var customer []models.Customer

	err := database.DB.First(&customer, "ID = ? ", customerID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    customer,
	})
}

// func CustomerUpdate(c *fiber.Ctx) error {

// }

func AddressDelete(c *fiber.Ctx) error {
	customerID := c.Params("ID")

	var customer []models.Customer

	err := database.DB.Debug().First(&customer, "ID = ?", customerID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&customer).Error

	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "address was deleted",
	})

}
