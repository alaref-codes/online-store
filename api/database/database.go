package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	DBconn *gorm.DB
)

type Spec struct {
	Password string
}

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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DBconn = db
	if err != nil {
		log.Fatal(err.Error())
	}

	DBconn.AutoMigrate(&Product{})
	fmt.Println("Database connection set")
}
