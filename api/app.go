package main

import (
	"log"

	"github.com/alaref-codes/onlin-store/api/database"
	"github.com/alaref-codes/onlin-store/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {

	// Products routes
	app.Get("/products", routes.GetProducts)
	app.Get("/products/:id", routes.GetOneProducts)
	app.Post("/products", routes.PostProducts)
	app.Delete("/products/:id", routes.DeleteProducts)

	// Order routes
	app.Get("/Orders", routes.GetOrders)
	app.Get("/Orders/:id", routes.GetOneOrder)
	app.Post("/Orders", routes.PostOrder)
	app.Delete("/Orders/:id", routes.DeleteOrder)

	// Restricted routes

}

func main() {

	// Initilizing the fiber instance
	app := fiber.New()

	// setting up the routes of the api
	setupRoutes(app)

	// Initilizing the mysql database connection
	database.InitDatabase()

	app.Use(logger.New())
	log.Fatal(app.Listen(":3000"))

}
