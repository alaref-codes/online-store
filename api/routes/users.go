package routes

import (
	"fmt"
	"regexp"
	"time"

	"github.com/alaref-codes/onlin-store/api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DBconn
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	var user database.User
	err := c.BodyParser(&user)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	result := db.Where("email = ?", user.Email).Find(&user)
	if result.RowsAffected != 0 {
		return fiber.NewError(503, "Record already exists")
	}
	if !emailRegex.MatchString(user.Email) {
		return fiber.NewError(503, "Not a valid email")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user = database.User{
		Email:    user.Email,
		Password: string(hash),
	}
	db.Create(&user)

	return c.JSON(fiber.Map{
		"message":      "User created successfully",
		"Created user": user,
	})
}

func Signin(c *fiber.Ctx) error {
	db := database.DBconn
	var user database.User
	var newUser database.User
	err := c.BodyParser(&newUser)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	result := db.Where("email = ?", newUser.Email).Find(&user)
	if result.RowsAffected == 0 {
		return fiber.NewError(404, "Auth failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password))
	if err != nil {
		return fiber.NewError(401, "Auth failed")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	// claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token":   t,
		"message": "Auth success",
	})

	// return c.SendString("Auth success")user
}

// Getting users : complete
func GetUsers(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	db := database.DBconn
	var users []database.User
	db.Find(&users)
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Welcome user %s", email),
		"users":   users,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DBconn
	email := c.Params("email")
	var user database.User
	result := db.Where("email = ?", email).Find(&user)

	if result.RowsAffected == 0 {
		return fiber.NewError(404, "Record doesn't exists")
	}

	result.Delete(&user)
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
		"users":   user,
	})
}
