package models

import "car_scraper/database"

type UserRepository struct{}

func (u UserRepository) GetUserByEmail(email string) User {
	var user User
	database.DB.Model(&User{}).First(&user, "email = ?", email)

	return user
}

func (u UserRepository) GetUserById(id uint8) User {
	var user User
	database.DB.Model(&User{}).Preload("Filters").First(&user, id)

	return user
}

type FilterRepository struct {}

func (r FilterRepository) SaveFilter(filter *Filter) error {
	result := database.DB.Model(&Filter{}).Create(filter)

	return result.Error
}