package models

import (
	"car_scraper/database"
	"golang.org/x/crypto/bcrypt"
)

func InitiateModels() {
	database.DB.AutoMigrate(&User{})
	database.DB.AutoMigrate(&Filter{})
	database.DB.AutoMigrate(&Car{})
}

type User struct {
	ID       uint8    `gorm:"primaryKey;autoincrement;not null"`
	Email    string   `gorm:"type:varchar(124);not null"`
	Password string   `gorm:"type:text;not null"`
	Filters  []Filter `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt int
	UpdatedAt int
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

type Filter struct {
	ID     uint32 `gorm:"primaryKey;autoincrement;not null" json:"id"`
	UserID uint8
	User   User
	Type   string `gorm:"type:varchar(32);not null" json:"type"`
	Search string `gorm:"type:varchar(450);not null" json:"search"`
	Cars []Car
	CreatedAt int
	UpdatedAt int
}

func (f *Filter) GetCarLinks() []string {
	result := make([]string, 0)

	for _, car := range f.Cars {
		result = append(result, car.Link)
	}

	return result
}

type Car struct {
	ID       uint64 `gorm:"primaryKey;autoincrement;not null"`
	FilterID uint32 `gorm:"not null"`
	Filter   Filter `gorm:"foreignKey:FilterID;not null;constraints:OnDelete:CASCADE"`
	Link     string `gorm:"type:varchar(150)"`
}

type CarDTO struct {
	ID          string
	Link		string
	Title       string
	Image       string
	Description string
	Price       string
	TopOffer    bool
}
