package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicErrorHandle(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": err.Error(),
	})
}