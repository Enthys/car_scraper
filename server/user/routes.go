package user

import (
	"car_scraper/database"
	"car_scraper/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserJSON struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var requestData RegisterUserJSON

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestData.Password != requestData.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	var userModel = database.DB.Model(&models.User{})

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatalln("Failed to encypt password")
		return
	}

	userModel.Create(&models.User{
		Email:    requestData.Email,
		Password: string(passwordHash),
	})

	c.JSON(http.StatusNoContent, gin.H{})
}
