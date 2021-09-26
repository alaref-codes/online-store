package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	return c.SendString("All products routes")
}

func GetOneProducts(c *fiber.Ctx) error {
	return c.SendString("One product routes")
}

func PostProducts(c *fiber.Ctx) error {
	return c.SendString("Post products routes")
}

func DeleteProducts(c *fiber.Ctx) error {
	return c.SendString("Delete products routes")
}
