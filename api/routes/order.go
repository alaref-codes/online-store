package routes

import (
	"fmt"
	"strconv"

	"github.com/alaref-codes/onlin-store/api/database"
	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	db := database.DBconn
	var orders []database.Order
	db.Find(&orders)
	return c.JSON(fiber.Map{
		"Message": "All Orders routes",
		"Orders":  orders,
	})
}

func GetOneOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn
	var orders database.Order
	db.Preload("Products").Find(&orders, id)
	return c.JSON(orders)
}

func PostOrder(c *fiber.Ctx) error {

	orderMap := make(map[string]int)

	c.BodyParser(&orderMap)

	var order database.Order

	for id, qty := range orderMap {
		p := database.Product{}

		id, _ := strconv.ParseUint(id, 10, 64)

		p.ID = uint(id)

		p.Quantity = qty

		order.Products = append(order.Products, p)
	}

	db := database.DBconn
	// var order database.Order
	// err := c.BodyParser(&order)

	// if err != nil {
	// 	return fiber.NewError(401, "Something went wrong")
	// }
	fmt.Printf("%v", order)
	if err := db.Save(&order).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"Message": "Order created successfully",
		"Order":   order,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn
	var order database.Order

	// db.Where("id = ?", id).Delete(&sub)  This one works too

	result := db.First(&order, id)

	if result.RowsAffected == 0 {
		return fiber.NewError(503, "Record doesn't exists")
	}

	result.Delete(&order)

	return c.SendString("Order Deleted successfully")
}
