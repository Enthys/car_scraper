package scraper

import (
	"car_scraper/models"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/elliotchance/orderedmap"
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

type PageSearchOptions interface {
	GetSearchType() string
}

type ICarCollection interface {
	AddCars(cars *orderedmap.OrderedMap)
	AddNewCars(seenCars, newCars map[string]models.CarDTO)
}

type CarCollection struct {
	ICarCollection
	Cars            map[string]models.CarDTO
	SeenTopOfferCar bool
	SeenNormalCar   bool
}

type Retriever interface {
	ParseSearchOptionsToValues(searchOptions PageSearchOptions) url.Values
	GetSearchResults(options PageSearchOptions) (string, string)
	GetCars(search PageSearchOptions, collection ICarCollection, page int) ICarCollection
	GetNewCars(search PageSearchOptions, page int) *orderedmap.OrderedMap
}

func GetScraper(scraperType string) (Scraper, error) {
	switch scraperType {
	case FilterTypeMobileBgCar:
		return MobileBGScraper{}, nil
	case FilterTypeMobileBgBike:
		return MobileBGScraper{}, nil
	case FilterTypeMobileBgBus:
		return MobileBGScraper{}, nil
	case FilterTypeCarsBgCar:
		return CarsBGScraper{}, nil
	case FilterTypeCarsBgBus:
		return CarsBGScraper{}, nil
	case FilterTypeCarsBgBike:
		return CarsBGScraper{}, nil
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


