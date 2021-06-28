package models

import (
	"car_scraper/database"
	"fmt"
)

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

func (u UserRepository) GetUsers() []User {
	var users []User
	database.DB.
		Preload("Filters").
		Find(&users)

	return users
}

type FilterRepository struct {}

func (r *FilterRepository) SaveFilter(filter *Filter) error {
	result := database.DB.Model(&Filter{}).Create(&filter)

	return result.Error
}

func (r *FilterRepository) UpdateFilter(filter *Filter) error {
	result := database.DB.Model(filter).Updates(&filter)

	return result.Error
}

func (r *FilterRepository) GetFilterByID(id uint32) Filter {
	var filter Filter
	database.DB.Preload("Cars").First(&filter, id)

	return filter
}

func (r *FilterRepository) DeleteFilter(f *Filter) error {
	fmt.Printf("%+v", f)
	result := database.DB.Model(f).Delete(f)

	return result.Error
}

func (r *FilterRepository) DeleteOldCars(c *CarRepository, f *Filter) error {
	if len(f.Cars) < 50 {
		return nil
	}

	shouldDelete := false

	for _, car := range f.Cars {
		if shouldDelete {
			if err := c.DeleteCar(&car); err != nil {
				println("Failed to delete car: ", car.ID)
			}
		}
		shouldDelete = !shouldDelete
	}

	return nil
}

type CarRepository struct {}

func (r CarRepository) SaveCar(car *Car) error {
	result := database.DB.Model(&Car{}).Create(&car)

	return result.Error
}

func (r CarRepository) DeleteCar(c *Car) error {
	result := database.DB.Model(c).Delete(c)

	return result.Error
}
