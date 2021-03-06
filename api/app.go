package main

import (
	"log"

	"github.com/alaref-codes/onlin-store/api/database"
	"github.com/alaref-codes/onlin-store/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func setupRoutes(app *fiber.App) {

	// Products routes
	app.Get("/products", routes.GetProducts)
	app.Get("/products/:id", routes.GetOneProducts)
	app.Post("/products", routes.PostProducts)
	app.Delete("/products/:id", routes.DeleteProducts)

	app.Post("/users/signin", routes.Signin)
	app.Post("/users/signup", routes.CreateUser) // Sign up page

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	// Restricted routes
	app.Get("/orders", routes.GetOrders)
	app.Get("/orders/:id", routes.GetOneOrder)
	app.Post("/orders", routes.PostOrder)
	app.Delete("/orders/:id", routes.DeleteOrder)

	app.Get("/users", routes.GetUsers)

	app.Delete("/users/:email", routes.DeleteUser)

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
