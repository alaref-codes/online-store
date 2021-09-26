package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	return c.SendString("All Orders routes")
}

func GetOneOrder(c *fiber.Ctx) error {
	return c.SendString("One Order routes")
}

func PostOrder(c *fiber.Ctx) error {
	return c.SendString("Post Orders routes")
}

func DeleteOrder(c *fiber.Ctx) error {
	return c.SendString("Delete Orders routes")
}
