package scraper

import (
	"car_scraper/models"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/elliotchance/orderedmap"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type CarsBGPageSearchOptions struct {
	PageSearchOptions
	Type       string `json:"type"`
	TypeOffer  string `cars_bg:"typeOffer"`
	Page       string `json:"page,omitempty" cars_bg:"page"`
	Brand      string `json:"brandId"        cars_bg:"brandId"`
	Model      string `json:"model"          cars_bg:"models[]"`
	YearStart  string `json:"yearFrom"       cars_bg:"yearFrom"`
	YearEnd    string `json:"yearTo"         cars_bg:"yearTo"`
	PriceStart string `json:"priceFrom"      cars_bg:"priceFrom"`
	PriceEnd   string `json:"priceTo"        cars_bg:"priceTo"`
}

func (c CarsBGPageSearchOptions) GetSearchType() string {
	return c.Type
}

type CarsBGBikePageSearchOptions struct {
	CarsBGPageSearchOptions
	Type       string `json:"type"`
	TypeOffer  string `cars_bg:"typeOffer"`
	Page       string `json:"page,omitempty" cars_bg:"page"`
	Brand      string `json:"brandId"        cars_bg:"brandId"`
	Model      string `json:"model"          cars_bg:"model_moto"`
	YearStart  string `json:"yearFrom"       cars_bg:"yearFrom"`
	YearEnd    string `json:"yearTo"         cars_bg:"yearTo"`
	PriceStart string `json:"priceFrom"      cars_bg:"priceFrom"`
	PriceEnd   string `json:"priceTo"        cars_bg:"priceTo"`
}

func (c CarsBGBikePageSearchOptions) GetSearchType() string {
	return c.Type
}

type CarsBGScraper struct {
	Scraper
}

func (c CarsBGScraper) CreateFilterFromFilterArgsString(filterType, filterArgs string) (*models.Filter, error) {
	var args CarsBGPageSearchOptions
	err := json.Unmarshal([]byte(filterArgs), &args)
	if err != nil {
		return nil, err
	}

	filter := &models.Filter{
		Search: filterArgs,
		Type:   filterType,
	}

	return filter, nil
}

func (c CarsBGScraper) InitiateFilter(filter *models.Filter) error {
	collection := carsBGCollection{
		CarCollection{
			Cars:            make(map[string]models.CarDTO),
			SeenNormalCar:   false,
			SeenTopOfferCar: false,
		},
	}

	retriever := carsBGRetriever{}

	searchParams, err := GetSearchParams(filter)
	if err != nil {
		return err
	}

	collection = retriever.GetCars(searchParams, collection, 1).(carsBGCollection)
	for _, car := range collection.Cars {
		filter.Cars = append(filter.Cars, models.Car{
			FilterID: filter.ID,
			Link:     car.Link,
		})
	}

	return nil
}

func getSearchOptions(searchParams PageSearchOptions) PageSearchOptions {
	switch searchParams.GetSearchType() {
	case FilterTypeCarsBgBike:
		return searchParams.(CarsBGBikePageSearchOptions)
	default:
		return searchParams.(CarsBGPageSearchOptions)
	}
}

func (c CarsBGScraper) GetNewCars(filter *models.Filter) *orderedmap.OrderedMap {
	oldCarLinks := filter.GetCarLinks()

	for i, link := range oldCarLinks {
		println(i, link)
	}

	searchParams, err := GetSearchParams(filter)

	if err != nil {
		panic(err)
	}
	
	carsBGSearchParams := getSearchOptions(searchParams)
	retriever := carsBGRetriever{}
	newCars := orderedmap.NewOrderedMap()

	page := 1
	seenOldCar := false
	for !seenOldCar {
		println("Checking page: ", page)
		cars := retriever.GetNewCars(carsBGSearchParams, page)

		for _, key := range cars.Keys() {
			carVal, _ := cars.Get(key)
			car := carVal.(models.CarDTO)

			println(car.Link)

			if InSliceString(oldCarLinks, car.Link) {
				seenOldCar = true
				break
			}

			newCars.Set(car.Link, car)
		}

		if seenOldCar == true || cars.Len() == 0 {
			break
		}

		log.Printf("New Cars %v", len(newCars.Keys()))
		page += 1
	}

	return newCars
}

type carsBGRetriever struct {
	Retriever
}

func (c carsBGRetriever) GetCars(search PageSearchOptions, collection ICarCollection, page int) ICarCollection {
	searchResults, _ := c.GetSearchResults(search)
	decoder := carsBGDecoder{}
	cars := decoder.GetCarsFromPageResults(searchResults)
	collection.AddCars(cars)

	return collection
}

func (c carsBGRetriever) GetNewCars(search PageSearchOptions, page int) *orderedmap.OrderedMap {
	var carsBGSearchOptions CarsBGPageSearchOptions
	carsBGSearchOptions = search.(CarsBGPageSearchOptions)
	carsBGSearchOptions.Page = strconv.Itoa(page)

	searchResults, _ := c.GetSearchResults(carsBGSearchOptions)
	decoder := carsBGDecoder{}

	return decoder.GetCarsFromPageResults(searchResults)
}

func (c carsBGRetriever) ParseSearchOptionsToValues(searchOptions PageSearchOptions) url.Values {
	valueReflection := reflect.ValueOf(searchOptions)
	typeReflection := reflect.TypeOf(searchOptions)
	var values = url.Values{}
	for i := 0; i < typeReflection.NumField(); i++ {
		searchField, found := typeReflection.Field(i).Tag.Lookup("cars_bg")
		if !found {
			continue
		}

		searchValue := reflect.Indirect(valueReflection).Field(i).String()
		if searchValue == "" {
			continue
		}

		values.Add(searchField, searchValue)
	}

	return values
}

func getOptionsType(options PageSearchOptions) string {
	if _, ok := options.(CarsBGPageSearchOptions); ok {
		return options.(CarsBGPageSearchOptions).Type
	}
	if _, ok := options.(CarsBGBikePageSearchOptions); ok {
		return options.(CarsBGBikePageSearchOptions).Type
	}

	panic("Invalid options");
}

func (c carsBGRetriever) GetSearchResults(options PageSearchOptions) (string, string) {
	client := http.Client{}
	var url string

	switch getOptionsType(options) {
	case "car":
		url = "https://www.Cars.bg/carslist.php"
		break
	case "bus":
		url = "https://www.Cars.bg/buseslist.php"
		break
	case "bike":
		url = "https://www.Cars.bg/motolist.php"
		break
	default:
		panic("Invalid CarsBG search option.")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	query := req.URL.Query()
	queryValues := c.ParseSearchOptionsToValues(options)
	for key, keyValues := range queryValues {
		for _, value := range keyValues {
			query.Set(key, value)
		}
	}

	req.URL.RawQuery = query.Encode()

	println(req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body), ""
}

type carsBGDecoder struct {
	Decoder
}

func (c carsBGDecoder) GetOfferTitle(doc *goquery.Document) string {
	return doc.
		Find("div.card__primary > h5").
		First().
		Text()
}

func (c carsBGDecoder) IsTopOffer(doc *goquery.Document) bool {
	return false
}

func (c carsBGDecoder) GetOfferDescription(doc *goquery.Document) string {
	partOne := doc.Find("div.card__secondary.mdc-typography.mdc-typography--body1.black").First().Text()
	partTwo := doc.Find("div.card__secondary.mdc-typography.mdc-typography--body2").First().Text()

	return fmt.Sprintf("%s\n%s", partOne, partTwo)
}

func (c carsBGDecoder) GetOfferPrice(doc *goquery.Document) string {
	return doc.Find("div.mdc-card__media.mdc-card__media--16-9 > div > h6").First().Text()
}

func (c carsBGDecoder) GetOfferImage(doc *goquery.Document) string {
	style, exists := doc.
		Find("a").First().
		Find("div").First().
		Attr("style")
	if !exists {
		return ""
	}

	style = strings.Replace(style, "background-image: url(\"", "", 1)
	style = strings.Replace(style, "\");", "", 1)

	return style
}

func (c carsBGDecoder) GetOfferID(doc *goquery.Document) string {
	link, _ := doc.Find("a").First().Attr("href")
	splitLink := strings.Split(link, "/")

	return splitLink[len(splitLink)-1]
}

func (c carsBGDecoder) GetOfferLink(doc *goquery.Document) string {
	link, _ := doc.Find("a").First().Attr("href")

	return link
}

func (c carsBGDecoder) GetCarsFromPageResults(pageResults string) *orderedmap.OrderedMap {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(pageResults))
	result := orderedmap.NewOrderedMap()

	doc.Find("#listContainer > div > div > div > div.mdc-layout-grid__inner.white-background.pageContainer.page-1 > div").
		Each(func(i int, carTable *goquery.Selection) {
			carDoc := goquery.NewDocumentFromNode(carTable.Nodes[0])
			title := c.GetOfferTitle(carDoc)

			if len(title) == 0 {
				return
			}

			carId := c.GetOfferID(carDoc)
			result.Set(carId, models.CarDTO{
				ID:          c.GetOfferID(carDoc),
				Link:        c.GetOfferLink(carDoc),
				Title:       c.GetOfferTitle(carDoc),
				Image:       c.GetOfferImage(carDoc),
				Description: c.GetOfferDescription(carDoc),
				Price:       c.GetOfferPrice(carDoc),
				TopOffer:    c.IsTopOffer(carDoc),
			})
		})

	return result
}

type carsBGCollection struct {
	CarCollection
}

func (c carsBGCollection) AddCars(cars *orderedmap.OrderedMap) {
	for _, key := range cars.Keys() {
		carVal, _ := cars.Get(key)
		car := carVal.(models.CarDTO)
		c.Cars[car.ID] = car

		if len(c.Cars) >= InitialCarLimit {
			break
		}
	}
}

func (c carsBGCollection) AddNewCars(seenCars, newCars map[string]models.CarDTO) {
	for _, newCar := range newCars {
		if _, ok := seenCars[newCar.ID]; ok {
			c.SeenNormalCar = true
			break
		}

		c.Cars[newCar.ID] = newCar
	}
}
