package main

import (
	"log"

	"github.com/alaref-codes/onlin-store/api/database"
	"github.com/alaref-codes/onlin-store/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
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
	viper.SetConfigFile(".env") // Reading the environment variable file
	viper.ReadInConfig()
	app := fiber.New()
	setupRoutes(app)
	database.InitDatabase()
	app.Use(logger.New())
	log.Fatal(app.Listen(":3000"))

}
