package scraper

import (
	"car_scraper/models"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/elliotchance/orderedmap"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

type MobileBGPageSearchOptions struct {
	PageSearchOptions
	Page        string `json:"f1,omitempty"      mobile_bg:"f1"`
	VehicleType string `json:"pubtype,omitempty" mobile_bg:"pubtype"`
	Brand       string `json:"f5"      mobile_bg:"f5"`
	Model       string `json:"f6"      mobile_bg:"f6"`
	YearStart   string `json:"f10"     mobile_bg:"f10"`
	YearEnd     string `json:"f11"     mobile_bg:"f11"`
	PriceStart  string `json:"f7"      mobile_bg:"f7"`
	PriceEnd    string `json:"f8"      mobile_bg:"f8"`
	Currency    string `json:"f9"      mobile_bg:"f9"`
	Slink 		string
}

func (m *MobileBGPageSearchOptions) SetSlink(slink string) {
	m.Slink = slink
}

type MobileBGScraper struct {
	Scraper
}

func (m MobileBGScraper) CreateFilterFromFilterArgsString(filterType string, filterArgs string) (*models.Filter, error) {
	var args MobileBGPageSearchOptions
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

func (m MobileBGScraper) InitiateFilter(filter *models.Filter) error {
	collection := mobileBGCollection{
		CarCollection{
			Cars:            make(map[string]models.CarDTO),
			SeenTopOfferCar: false,
			SeenNormalCar:   false,
		},
		make(map[string]models.CarDTO),
		make(map[string]models.CarDTO),
	}
	retriever := mobileBGRetriever{}

	searchParams, err := GetSearchParams(filter)
	if err != nil {
		return err
	}

	collection = retriever.GetCars(searchParams, collection, 1).(mobileBGCollection)
	for _, car := range collection.topOfferCars {
		filter.Cars = append(filter.Cars, models.Car{
			FilterID: filter.ID,
			Link:     car.Link,
		})
	}

	for _, car := range collection.normalCars {
		filter.Cars = append(filter.Cars, models.Car{
			FilterID: filter.ID,
			Link:     car.Link,
		})
	}

	return nil
}

func clearSlinkFromCarLink(carLink string) string {
	slinkPattern := regexp.MustCompile("^(.*?)&slink=.+")
	return slinkPattern.ReplaceAllString(carLink, "${1}")
}

func getFilterCarLinks(filter *models.Filter) []string {
	linksWithoutSlink := make([]string, 0)
	oldCarLinks := filter.GetCarLinks()

	for _, link := range oldCarLinks {
		linksWithoutSlink = append(linksWithoutSlink, clearSlinkFromCarLink(link))
	}

	return linksWithoutSlink
}

func (m MobileBGScraper) GetNewCars(filter *models.Filter) *orderedmap.OrderedMap {
	oldCarLinks := getFilterCarLinks(filter)

	searchParams, err := GetSearchParams(filter)
	if err != nil {
		panic(err)
	}

	mobileBGSearchParams := searchParams.(MobileBGPageSearchOptions)
	retriever := mobileBGRetriever{}
	newCars := orderedmap.NewOrderedMap()
	seenOldTopOfferCar := false
	seenOldNormalCar := false

	page := 1
	for {
		cars := retriever.GetNewCars(&mobileBGSearchParams, page)

		for _, key := range cars.Keys() {
			carVal, _ := cars.Get(key)
			car := carVal.(models.CarDTO)

			if InSliceString(oldCarLinks, clearSlinkFromCarLink(car.Link)) {
				if car.TopOffer {
					seenOldTopOfferCar = true
				} else {
					seenOldNormalCar = true
				}
				if seenOldTopOfferCar == true && seenOldNormalCar == true {
					break
				}

				continue
			}

			if seenOldTopOfferCar && car.TopOffer {
				continue
			}

			newCars.Set(car.Link, car)
		}

		if seenOldTopOfferCar == true && seenOldNormalCar == true || cars.Len() == 0 {
			break
		}

		page += 1
	}

	return newCars
}

type mobileBGRetriever struct {
	Retriever
}

func (r mobileBGRetriever) GetCars(
	search PageSearchOptions,
	collection ICarCollection,
	page int,
) ICarCollection {
	decoder := mobileBGDecoder{}
	slink := ""
	for {
		if len(collection.(mobileBGCollection).topOfferCars) == InitialCarLimit / 2 && len(collection.(mobileBGCollection).normalCars) == InitialCarLimit / 2 {
			break
		}

		if slink == "" {
			_, slink = r.GetSearchResults(search)
		}
		searchResult := r.getSearchBySlink(slink, page)
		cars := decoder.GetCarsFromPageResults(searchResult)
		if cars.Len() == 0 {
			break
		}
		collection.AddCars(cars)
		page += 1
	}

	return collection
}

func (r mobileBGRetriever) GetNewCars(
	search *MobileBGPageSearchOptions,
	page int,
) *orderedmap.OrderedMap {
	if search.Slink == "" {
		_, slink := r.GetSearchResults(*search)
		search.SetSlink(slink)
	}

	searchResult := r.getSearchBySlink(search.Slink, page)

	decoder := mobileBGDecoder{}
	return decoder.GetCarsFromPageResults(searchResult)
}

func (r mobileBGRetriever) ParseSearchOptionsToValues(searchOptions PageSearchOptions) url.Values {
	valueReflection := reflect.ValueOf(searchOptions)
	typeReflection := reflect.TypeOf(searchOptions)
	var values = url.Values{}
	values.Add("act", "3")
	for i := 0; i < typeReflection.NumField(); i++ {
		searchField, found := typeReflection.Field(i).Tag.Lookup("mobile_bg")
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

func (r mobileBGRetriever) GetSearchResults(options PageSearchOptions) (string, string) {
	myURL := "https://mobile.bg/pcgi/mobile.cgi"
	nextURL := myURL
	var i int
	var resp *http.Response
	for i < 100 {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}

		resp, err := client.Post(
			nextURL,
			"application/x-www-form-urlencoded",
			strings.NewReader(r.ParseSearchOptionsToValues(options).Encode()),
		)

		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == http.StatusOK {
			break
		} else {
			nextURL = resp.Header.Get("Location")
			if strings.Contains(nextURL, "slink") {
				break
			}
			i += 1
		}
	}

	resp, err := http.Get(nextURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pageContent, err := charmap.Windows1251.NewDecoder().String(string(body))
	if err != nil {
		log.Fatal(err)
	}

	slinkMatches := regexp.MustCompile(".+slink=(.+)&")
	slink := slinkMatches.FindAllStringSubmatch(nextURL, -1)[0][1]

	return pageContent, slink
}

func (r mobileBGRetriever) getSearchBySlink(slink string, page int) string {
	requestUrl := fmt.Sprintf("https://mobile.bg/pcgi/mobile.cgi?act=3&slink=%s&f1=%v", slink, page)

	resp, err := http.Get(requestUrl)
	if err != nil {
		counter := 0
		for counter < 5 || err != nil {
			resp, err = http.Get(requestUrl)
			counter += 1
		}
		if counter == 5 && err != nil {
			log.Fatal("Failed to retrieve MobileBG Search results")
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pageContent, err := charmap.Windows1251.NewDecoder().String(string(body))
	if err != nil {
		log.Fatal(err)
	}

	return pageContent
}

type mobileBGDecoder struct {
	Decoder
}

func (d mobileBGDecoder) GetOfferTitle(doc *goquery.Document) string {
	return doc.
		Find("tbody tr td.valgtop a.mmm").
		First().
		Text()
}

func (d mobileBGDecoder) IsTopOffer(doc *goquery.Document) bool {
	return doc.Find("tbody tr td.algright img.noborder").Length() == 3
}

func (d mobileBGDecoder) GetOfferDescription(doc *goquery.Document) string {
	return doc.
		Find("tbody tr:nth-child(3) td").
		First().
		Text()
}

func (d mobileBGDecoder) GetOfferPrice(doc *goquery.Document) string {
	return doc.
		Find("tbody tr td.algright.valgtop span.price").
		First().
		Text()
}

func (d mobileBGDecoder) GetOfferImage(doc *goquery.Document) string {
	link, _ := doc.
		Find("tbody tr td.algcent.valgmid a.photoLink img.noborder").
		First().
		Attr("src")

	return link
}

func (d mobileBGDecoder) GetOfferID(doc *goquery.Document) string {
	link, _ := doc.
		Find("tbody tr td.algcent.valgmid a.photoLink").
		First().
		Attr("href")

	linkUrl, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}

	return linkUrl.Query().Get("adv")
}

func (d mobileBGDecoder) GetOfferLink(doc *goquery.Document) string {
	link, _ := doc.
		Find("tbody tr td.algcent.valgmid a.photoLink").
		First().
		Attr("href")

	return link[2:]
}

func (d mobileBGDecoder) GetCarsFromPageResults(pageResults string) *orderedmap.OrderedMap {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(pageResults))
	result := orderedmap.NewOrderedMap()
	doc.Find("form[name=\"search\"]").
		First().
		Find("table").
		Each(func(i int, carTable *goquery.Selection) {
			carDoc := goquery.NewDocumentFromNode(carTable.Nodes[0])
			title := d.GetOfferTitle(carDoc)

			if len(title) == 0 {
				return
			}

			carId := d.GetOfferID(carDoc)
			result.Set(carId, models.CarDTO{
				ID:          d.GetOfferID(carDoc),
				Link:        d.GetOfferLink(carDoc),
				Title:       d.GetOfferTitle(carDoc),
				Image:       d.GetOfferImage(carDoc),
				Description: d.GetOfferDescription(carDoc),
				Price:       d.GetOfferPrice(carDoc),
				TopOffer:    d.IsTopOffer(carDoc),
			})
		})

	return result
}

type mobileBGCollection struct {
	CarCollection
	topOfferCars, normalCars map[string]models.CarDTO
}

func (col mobileBGCollection) AddCars(cars *orderedmap.OrderedMap) {
	for _, key := range cars.Keys() {
		carVal, _ := cars.Get(key)
		car := carVal.(models.CarDTO)
		if car.TopOffer && len(col.topOfferCars) != InitialCarLimit / 2 {
			col.topOfferCars[car.ID] = car
			col.Cars[car.ID] = car
		} else if !car.TopOffer && len(col.normalCars) != InitialCarLimit / 2 {
			col.normalCars[car.ID] = car
			col.Cars[car.ID] = car
		}

		if len(col.Cars) >= InitialCarLimit {
			break
		}
	}
}

func (col mobileBGCollection) AddNewCars(seenCars, newCars map[string]models.CarDTO) {
	for _, newCar := range newCars {
		if _, ok := seenCars[newCar.ID]; ok {
			if newCar.TopOffer {
				col.SeenTopOfferCar = true
				continue
			}
			col.SeenNormalCar = true
			break
		}
		if newCar.TopOffer && col.SeenTopOfferCar {
			continue
		}

		col.Cars[newCar.ID] = newCar
	}
}
