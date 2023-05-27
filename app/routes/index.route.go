package routes

import (
	"app-transaction/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteFunc(r *fiber.App) {

	// // transaction
	// r.Get("/transactions", controllers.TransactionGetAll)
	// r.Post("/transaction", controllers.CreateTransaction)
	// r.Put("/transaction/:id", controllers.CreateTransaction)
	// r.Delete("/transaction/:id", controllers.DeleteTransaction)

	// // payment Method
	r.Get("/payments", controllers.PaymentGetAll)
	r.Get("/payment/:id", controllers.PaymentGetByID)
	r.Post("/payment", controllers.PaymentCreate)
	// r.Put("/payment/:id", controllers.PaymentUpdate)
	r.Delete("/paymnet/:id", controllers.PaymentDelete)

	// // product
	r.Get("/products", controllers.ProductGetAll)
	r.Get("/product/:id", controllers.ProductGetByID)
	r.Post("/product", controllers.ProductCreate)
	// r.Put("/product/:id", controllers.ProductUpdate)
	r.Delete("/product/:id", controllers.ProductDelete)

	// // customer
	r.Get("/customers", controllers.CustomerGetAll)
	r.Get("/customers/:id", controllers.CustomerGetByID)
	r.Post("/customer", controllers.CustomerCreate)
	// r.Put("/customer/:id", controllers.CustomerUpdate)
	r.Delete("/customer/:id", controllers.CustomerDelete)

	// // address
	r.Get("/addresss", controllers.AddressGetAll)
	r.Get("/address/:id", controllers.AddressGetByID)
	r.Post("/address", controllers.AddressCreate)
	// r.Put("/address/:id", controllers.AddressUpdate)
	r.Delete("/address/:id", controllers.AddressDelete)
}
