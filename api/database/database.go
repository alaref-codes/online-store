package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Order struct {
	ID       uint      `gorm:"primaryKey"`
	Products []Product `gorm:"many2many:order_products;"`
}

type Product struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Quantity int
}

var (
	DBconn *gorm.DB
)

func InitDatabase() {
	var err error
	viper.SetConfigName("config") // Reading the environment variable file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	pass := viper.Get("PASSWORD").(string)

	if err != nil {
		log.Fatal(err.Error())
	}

	dbName := "online_store"
	dsn := "alaref:" + pass + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Info,
	})
	DBconn = db
	if err != nil {
		log.Fatal(err.Error())
	}

	DBconn.AutoMigrate(&Order{}, &User{})
	fmt.Println("Database Connection Set")
}
