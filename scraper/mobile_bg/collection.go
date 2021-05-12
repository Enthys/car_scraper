package mobile_bg

import "car_scraper/models"

const InitialCarLimit = 10

type MobileBGCollection struct {
	cars            map[string]models.CarDTO
	topOfferCars    map[string]models.CarDTO
	normalCars      map[string]models.CarDTO
	seenTopOfferCar bool
	seenNormalCar   bool
}

func (col MobileBGCollection) addCars(cars []models.CarDTO) {
	for _, car := range cars {
		if car.TopOffer && len(col.topOfferCars) != InitialCarLimit {
			col.topOfferCars[car.ID] = car
			col.cars[car.ID] = car
		} else if !car.TopOffer && len(col.normalCars) != InitialCarLimit {
			col.normalCars[car.ID] = car
			col.cars[car.ID] = car
		}

		if len(col.normalCars) >= InitialCarLimit {
			break
		}
	}
}

func (col MobileBGCollection) addNewCars(seenCars, newCars map[string]models.CarDTO) {
	for _, newCar := range newCars {
		if _, ok := seenCars[newCar.ID]; ok {
			if newCar.TopOffer {
				col.seenTopOfferCar = true
				continue
			}
			col.seenNormalCar = true
			break
		}
		if newCar.TopOffer && col.seenTopOfferCar {
			continue
		}

		col.cars[newCar.ID] = newCar
	}
}
