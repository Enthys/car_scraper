package filter

import (
	"car_scraper/models"
	"car_scraper/server/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCarFilters(c *gin.Context) {
	val, err := c.Get(middleware.UserId)
	if err != true {
		log.Fatal(err)
	}
	user := models.UserRepository{}.GetUserById(val.(uint8))

	c.JSON(http.StatusOK, user.Filters)
}

func CreateCarFilter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "CreateCarFilter",
	})
}

func DeleteCarFilter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetCarFilters",
	})
}
