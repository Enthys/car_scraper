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
	database.DB.Model(&User{}).First(&user, "id = ?", id)

	return user
}