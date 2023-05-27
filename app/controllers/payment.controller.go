package controllers

import (
	"app-transaction/database"
	"app-transaction/models"

	"github.com/gofiber/fiber/v2"
)

func PaymentCreate(c *fiber.Ctx) error {

	customer := new(models.Customer)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	database.DB.Create(&customer)

	return c.JSON(fiber.Map{
		"message": "create customer successfully",
		"post":    customer,
	})

}

func PaymentGetAll(c *fiber.Ctx) error {
	var customer []models.Customer

	database.DB.Preload("Customer").Find(&customer)

	return c.JSON(fiber.Map{
		"customer": customer,
	})
}

func PaymentGetByID(c *fiber.Ctx) error {
	paymentID := c.Params("ID")

	var payment []models.Payment

	err := database.DB.First(&payment, "ID = ? ", paymentID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    payment,
	})
}

// func CustomerUpdate(c *fiber.Ctx) error {

// }

func PaymentDelete(c *fiber.Ctx) error {
	paymentID := c.Params("ID")

	var payment []models.Customer

	err := database.DB.Debug().First(&payment, "ID = ?", paymentID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&payment).Error

	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "customer was deleted",
	})

}
