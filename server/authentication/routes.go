package authentication

import (
	"car_scraper/auth"
	"car_scraper/database"
	"car_scraper/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestJSON struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var requestData LoginRequestJSON

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid arguments passed.",
		})

		return
	}

	var user models.User
	database.DB.First(&user, "email = ?", requestData.Email)

	if err := user.CheckPassword(requestData.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Credentials",
		})

		return
	}

	jwtWrapper := auth.JwtWrapper{}

	token, err := jwtWrapper.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout",
	})
}
