package scraper

import (
	"car_scraper/models"
	"encoding/json"
	"github.com/elliotchance/orderedmap"
)

const InitialCarLimit = 10

type Scraper interface {
	CreateFilterFromFilterArgsString(filterType, filterArgs string) (*models.Filter, error)
	InitiateFilter(filter *models.Filter) error
	GetNewCars(filter *models.Filter) *orderedmap.OrderedMap
}

func GetSearchParams(f *models.Filter) (PageSearchOptions, error) {
	switch f.Type {
	case FilterTypeMobileBgCar:
		var searchOpts MobileBGPageSearchOptions
		err := json.Unmarshal([]byte(f.Search), &searchOpts)
		if err != nil {
			return nil, err
		}

		searchOpts.Slink = ""
		searchOpts.Page = "1"
		searchOpts.VehicleType = "1"
		return searchOpts, nil
	case FilterTypeMobileBgBike:
		var searchOpts MobileBGPageSearchOptions
		err := json.Unmarshal([]byte(f.Search), &searchOpts)
		if err != nil {
			return nil, err
		}

		searchOpts.Slink = ""
		searchOpts.Page = "1"
		searchOpts.VehicleType = "5"
		return searchOpts, nil
	case FilterTypeMobileBgBus:
		var searchOpts MobileBGPageSearchOptions
		err := json.Unmarshal([]byte(f.Search), &searchOpts)
		if err != nil {
			return nil, err
		}

		searchOpts.Slink = ""
		searchOpts.Page = "1"
		searchOpts.VehicleType = "3"
		return searchOpts, nil
	case FilterTypeCarsBgCar:
		var searchOpts CarsBGPageSearchOptions
		err := json.Unmarshal([]byte(f.Search), &searchOpts)
		if err != nil {
			return nil, err
		}

		searchOpts.Page = "1"
		searchOpts.TypeOffer = "1"
		return searchOpts, nil
	case FilterTypeCarsBgBus:
		var searchOpts CarsBGPageSearchOptions
		err := json.Unmarshal([]byte(f.Search), &searchOpts)
		if err != nil {
			return nil, err
		}

		searchOpts.Page = "1"
		searchOpts.TypeOffer = "2"
		return searchOpts, nil
	case FilterTypeCarsBgBike:
		var searchOpts CarsBGBikePageSearchOptions
		err := json.Unmarshal([]byte(f.Search), &searchOpts)
		if err != nil {
			return nil, err
		}

		searchOpts.Page = "1"
		searchOpts.TypeOffer = "3"
		return searchOpts, nil
	default:
		panic("Failed to create search params")
	}
}

func InSliceString(a []string, x string) bool {
	for _, str := range a {
		if str == x {
			return true
		}
	}

	return false
}