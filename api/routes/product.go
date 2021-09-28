package routes

import (
	"fmt"
	"strconv"

	"github.com/alaref-codes/onlin-store/api/database"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	db := database.DBconn
	var products []database.Product
	db.Find(&products)
	return c.JSON(fiber.Map{
		"Message":  "All products routes",
		"Products": products,
	})

}

func GetOneProducts(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn
	var product database.Product
	db.Find(&product, id)
	return c.JSON(product)

}

func PostProducts(c *fiber.Ctx) error {
	db := database.DBconn
	productMap := make(map[string]string)

	err := c.BodyParser(&productMap)

	if err != nil {
		return fiber.NewError(401, "Something went wrong")
	}

	var product database.Product

	for name, qty := range productMap {
		product.Name = name
		product.Quantity, _ = strconv.Atoi(qty)
	}
	fmt.Printf("%v", product)

	if err := db.Create(&product).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"Message": "Product created successfully",
		"Product": product,
	})
}

func DeleteProducts(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn
	var product database.Product

	result := db.First(&product, id)

	if result.RowsAffected == 0 {
		return fiber.NewError(503, "Record doesn't exists")
	}

	result.Delete(&product)

	return c.SendString("Product Deleted successfully")
}
