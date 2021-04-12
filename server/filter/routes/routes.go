package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCarFilters(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetCarFilters",
	})
}

func CreateCarFilrer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetCarFilters",
	})
}

func DeleteCarFilter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetCarFilters",
	})
}
