package main

import (
	"app-transaction/database"
	"app-transaction/database/migrations"
	"app-transaction/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// connect database

	database.DatabaseInit()

	// migration
	migrations.RunMigration()

	r := fiber.New()
	routes.RouteFunc(r)

	// cek route
	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Message": "hello-api",
			"kodeCek": "200",
		})
	})

	// PORT!
	r.Listen(":2345")

}
