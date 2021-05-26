package scraper

import "car_scraper/models"

type Scraper interface {
	CreateFilterFromString(filterArgs string) (*models.Filter, error)
	InitiateFilter(filter *models.Filter) error
}

type MobileBGScraper struct {
}

func (m MobileBGScraper) CreateFilterFromString(filterArgs string) (*models.Filter, error) {
	panic("implement me")
}

func (m MobileBGScraper) InitiateFilter(filter *models.Filter) error {
	panic("implement me")
}
