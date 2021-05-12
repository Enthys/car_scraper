package mobile_bg

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

type PageSearchOptions struct {
	SearchPage  string `json:"f1"      mobile_bg:"f1"`
	VehicleType string `json:"pubtype" mobile_bg:"pubtype"`
	Brand       string `json:"f5"      mobile_bg:"f5"`
	Model       string `json:"f6"      mobile_bg:"f6"`
	YearStart   string `json:"f10"     mobile_bg:"f10"`
	YearEnd     string `json:"f11"     mobile_bg:"f11"`
	PriceStart  string `json:"f7"      mobile_bg:"f7"`
	PriceEnd    string `json:"f8"      mobile_bg:"f8"`
	Currency    string `json:"f9"      mobile_bg:"f9"`
}

func ParseSearchOptionsToValues(searchOptions PageSearchOptions) url.Values {
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

func getSearchResults(options PageSearchOptions) (string, string) {
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
			strings.NewReader(ParseSearchOptionsToValues(options).Encode()),
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

func GetSearchBySlink(slink string, page int) string {
	myURL := fmt.Sprintf("https://mobile.bg/pcgi/mobile.cgi?act=3&slink=%s&f1=%v", slink, page)

	resp, err := http.Get(myURL)
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

	return pageContent
}
