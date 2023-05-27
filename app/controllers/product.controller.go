package controllers

import (
	"app-transaction/database"
	"app-transaction/models"

	"github.com/gofiber/fiber/v2"
)

func ProductCreate(c *fiber.Ctx) error {

	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	database.DB.Create(&product)

	return c.JSON(fiber.Map{
		"message": "create customer successfully",
		"post":    product,
	})

}

func ProductGetAll(c *fiber.Ctx) error {
	var product []models.Product

	err := database.DB.Preload("Product").Find(&product).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return c.JSON(fiber.Map{
		"product": product,
	})
}

func ProductGetByID(c *fiber.Ctx) error {
	productID := c.Params("ID")

	var product []models.Product

	err := database.DB.First(&product, "ID = ? ", productID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    product,
	})
}

// func CustomerUpdate(c *fiber.Ctx) error {

// }

func ProductDelete(c *fiber.Ctx) error {
	ProductID := c.Params("ID")

	var product []models.Product

	err := database.DB.Debug().First(&product, "ID = ?", ProductID).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&product).Error

	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "customer was deleted",
	})

}
