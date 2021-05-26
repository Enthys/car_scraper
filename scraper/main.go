package scraper

import (
	"car_scraper/models"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

const (
	FilterTypeMobileBgCar  = "MobileBGCar"
	FilterTypeMobileBgBike = "MobileBGBike"
	FilterTypeMobileBgBus  = "MobileBGBus"
	FilterTypeCarsBgCar    = "CarsBGCar"
	FilterTypeCarsBgBike   = "CarsBGBike"
	FilterTypeCarsBgBus    = "CarsBGBus"
)

type PageSearchOptions struct {
	SearchPage  string
}

type ICarCollection interface {
	AddCars(cars map[string]models.CarDTO)
	AddNewCars(seenCars, newCars map[string]models.CarDTO)
}

type CarCollection struct {
	ICarCollection
	cars            map[string]models.CarDTO
	seenTopOfferCar bool
	seenNormalCar   bool
}

type Retriever interface {
	ParseSearchOptionsToValues(searchOptions PageSearchOptions) url.Values
	GetSearchResults(options PageSearchOptions) (string, string)
	GetCars(search PageSearchOptions, collection CarCollection, page int) CarCollection
}

func GetScraper(scraperType string) (Scraper, error) {
	switch scraperType {
	case FilterTypeMobileBgCar:
		return MobileBGScraper{}, nil
	default:
		return nil, errors.New("invalid scraper type")
	}
}

type Decoder interface {
	GetOfferTitle(doc *goquery.Document) string
	IsTopOffer(doc *goquery.Document) bool
	GetOfferDescription(doc *goquery.Document) string
	GetOfferPrice(doc *goquery.Document) string
	GetOfferImage(doc *goquery.Document) string
	GetOfferID(doc *goquery.Document) string
	GetOfferLink(doc *goquery.Document) string
	GetCarsFromPageResults(pageResults string) map[string]models.CarDTO
}


