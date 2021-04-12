package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int8   `gorm:"primaryKey;autoincrement;not null"`
	Email    string `gorm:"type:varchar(124);not null"`
	Password string
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
