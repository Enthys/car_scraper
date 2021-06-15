package filter

import (
	"car_scraper/models"
	"car_scraper/scraper"
	"car_scraper/server/middleware"
	"car_scraper/server/utils"
	"io/ioutil"
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
	Type   string `json:"type"`
	Filter string `json:"filter"`
}

func CreateCarFilter(c *gin.Context) {
	var request CreateCarRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.BasicErrorHandle(c, err)
		return
	}

	scrapingService, err := scraper.GetScraper(request.Type)
	if err != nil {
		utils.BasicErrorHandle(c, err)
		return
	}

	filter, err := scrapingService.CreateFilterFromFilterArgsString(request.Type, request.Filter)
	if err != nil {
		utils.BasicErrorHandle(c, err)
		return
	}
	val, _ := c.Get(middleware.UserId)
	filter.UserID = val.(uint8)

	filterRepo := models.FilterRepository{}

	var fns []func(filter *models.Filter) error
	fns = append(
		fns,
		filterRepo.SaveFilter,
		scrapingService.InitiateFilter,
		filterRepo.UpdateFilter,
	)

	for _, fn := range fns {
		err := fn(filter)
		if err != nil {
			utils.BasicErrorHandle(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"requestData": request,
	})
}

func DeleteCarFilter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetCarFilters",
	})
}

func GetCarsBGBrandModels(c *gin.Context) {
	vehicleType := c.Param("type")
	brandId := c.Param("brandId")
	var url string
	switch vehicleType {
	case "car":
		url = "https://www.cars.bg/carmodel.php?brandId=" + brandId
		break
	case "bus":
		url = "https://www.cars.bg/carmodel.php?isBus=1&brandId=" + brandId
		break
	default:
		panic("Invalid carsbg vehicle type.")
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		panic(err)
	}

	c.Data(http.StatusOK, "text/plain", body)
}