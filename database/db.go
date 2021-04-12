package database

import (
	"log"

	"car_scraper/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := "root:secret_password@tcp(127.0.0.1:3306)/car_scraper"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
	log.Println("Connected to database!")
}

func InitiateModels() {
	DB.AutoMigrate(&models.User{})
}
