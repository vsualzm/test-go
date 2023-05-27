package controllers

import (
	"app-transaction/database"
	"app-transaction/models"

	"github.com/gofiber/fiber/v2"
)

func TransactionGetAll(c *fiber.Ctx) error {

	var transactions []models.Transaction

	database.DB.Preload("Customer").Preload("Product").Preload("payment").Find(&transactions)

	return c.JSON(fiber.Map{
		"transactions": transactions,
	})
}

func TransactionCreate(c *fiber.Ctx) error {

	transaction := new(models.Transaction)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	database.DB.Create(&transaction)

	return c.JSON(fiber.Map{
		"message": "create customer successfully",
		"post":    transaction,
	})

}

func TransactionGetByID(c *fiber.Ctx) error {
	transactionID := c.Params("ID")

	var transaction []models.Transaction

	err := database.DB.First(&transaction, "ID = ? ", transactionID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "transaction not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    transaction,
	})
}

func TransactionDelete(c *fiber.Ctx) error {
	TransactionID := c.Params("ID")

	var transaction []models.Transaction

	err := database.DB.Debug().First(&transaction, "ID = ?", TransactionID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "transaction not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&transaction).Error

	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "transaction was deleted",
	})

}
