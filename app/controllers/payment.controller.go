package controllers

import (
	"app-transaction/database"
	"app-transaction/models"

	"github.com/gofiber/fiber/v2"
)

func PaymentCreate(c *fiber.Ctx) error {

	payment := new(models.Payment)

	if err := c.BodyParser(payment); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	database.DB.Create(&payment)

	return c.JSON(fiber.Map{
		"message": "create customer successfully",
		"post":    payment,
	})

}

func PaymentGetAll(c *fiber.Ctx) error {
	var payment []models.Payment

	database.DB.Preload("Customer").Find(&payment)

	return c.JSON(fiber.Map{
		"customer": payment,
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

	var payment []models.Payment

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
