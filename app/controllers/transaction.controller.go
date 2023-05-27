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
