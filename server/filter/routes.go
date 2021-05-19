package filter

import (
	"car_scraper/models"
	"car_scraper/scraper"
	"car_scraper/server/middleware"
	"car_scraper/server/utils"
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

type CreateCarRequest struct {
	Type string `json:"type"`
	Filter string `json:"filter"`
}

func CreateCarFilter(c *gin.Context) {
	var request CreateCarRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.BasicErrorHandle(c, err)
		return
	}
	print(request.Filter)
	scrapingService, err := scraper.GetScrapingService(request.Type)
	if err != nil {
		utils.BasicErrorHandle(c, err)
		return
	}

	filter, err := scrapingService.CreateFilterFromString(request.Filter)
	if err != nil {
		utils.BasicErrorHandle(c, err)
		return
	}
	val, _ := c.Get(middleware.UserId)
	filter.UserID = val.(uint8)

	err = models.FilterRepository{}.SaveFilter(filter)
	if err != nil {
		utils.BasicErrorHandle(c, err)
		return
	}

	err = scrapingService.InitiateFilter(filter)
	if err != nil {
		utils.BasicErrorHandle(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"requestData": request,
	})
}

func DeleteCarFilter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetCarFilters",
	})
}
