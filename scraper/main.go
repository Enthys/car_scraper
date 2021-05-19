package scraper

import (
	"car_scraper/models"
	"car_scraper/scraper/mobile_bg"
	"encoding/json"
	"errors"
	"log"
)

const (
	FilterTypeMobileBgCar  = "MobileBGCar"
	FilterTypeMobileBgBike = "MobileBGBike"
	FilterTypeMobileBgBus  = "MobileBGBus"
	FilterTypeCarsBgCar    = "CarsBGCar"
	FilterTypeCarsBgBike   = "CarsBGBike"
	FilterTypeCarsBgBus    = "CarsBGBus"
)

type Scraper interface {
	CreateFilterFromString(filterArgs string) (*models.Filter, error)
	InitiateFilter(filter *models.Filter) error
}

type MobileBGScraper struct {}

func (s MobileBGScraper) CreateFilterFromString(filterArgs string) (*models.Filter, error) {
	var filterSearchArgs mobile_bg.PageSearchOptions
	err := json.Unmarshal([]byte(filterArgs), &filterSearchArgs)
	if err != nil {
		return nil, err
	}

	filterString, err := json.Marshal(filterSearchArgs)
	if err != nil {
		return nil, err
	}

	filter := models.Filter{
		Type:   "MobileBGCar",
		Search: string(filterString),
	}

	return &filter, nil
}

func (s MobileBGScraper) InitiateFilter(filter *models.Filter) error {
	log.Printf("%v", filter)

	return nil
}

func GetScrapingService(scraperType string) (Scraper, error) {
	switch scraperType {
	case FilterTypeMobileBgCar:
		return MobileBGScraper{}, nil
	default:
		return nil, errors.New("invalid scraper type")
	}
}
