package models

import (
	"car_scraper/database"
	"golang.org/x/crypto/bcrypt"
)

func InitiateModels() {
	database.DB.AutoMigrate(&User{})
	database.DB.AutoMigrate(&Filter{})
}

type User struct {
	ID       uint8    `gorm:"primaryKey;autoincrement;not null"`
	Email    string   `gorm:"type:varchar(124);not null"`
	Password string   `gorm:"type:text;not null"`
	Filters  []Filter `gorm:"foreignKey:UserID"`
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

type UserRepository struct{}

func (u UserRepository) GetUserByEmail(email string) User {
	var user User
	database.DB.Model(&User{}).First(&user, "email = ?", email)

	return user
}

func (u UserRepository) GetUserById(id uint8) User {
	var user User
	database.DB.Model(&User{}).First(&user, "id = ?", id)

	return user
}

type Filter struct {
	ID     uint32 `gorm:"primaryKey;autoincrement;not null"`
	UserID uint8  `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID;not null"`
	Type   string `gorm:"type:varchar(32);not null"`
	Search string `gorm:"type:varchar(300);not null"`
}

type Car struct {
	ID       uint64 `gorm:"primaryKey;autoincrement;not null"`
	FilterID uint32 `gorm:"not null"`
	Link     string `gorm:"type:varchar(150)"`
	CarDTO   CarDTO
}

type CarDTO struct {
	Title       string
	Image       string
	Description string
	Price       string
	TopOffer    bool
}
