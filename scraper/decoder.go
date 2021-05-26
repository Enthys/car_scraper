package scraper

import (
	"car_scraper/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"strings"
)

type MobileBGDecoder struct {
	Decoder
}

func (d MobileBGDecoder) GetOfferTitle(doc *goquery.Document) string {
	return doc.
		Find("tbody tr td.valgtop a.mmm").
		First().
		Text()
}

func (d MobileBGDecoder) IsTopOffer(doc *goquery.Document) bool {
	return doc.Find("tbody tr td.algright img.noborder").Length() == 3
}

func (d MobileBGDecoder) GetOfferDescription(doc *goquery.Document) string {
	return doc.
		Find("tbody tr:nth-child(3) td").
		First().
		Text()
}

func (d MobileBGDecoder) GetOfferPrice(doc *goquery.Document) string {
	return doc.
		Find("tbody tr td.algright.valgtop span.price").
		First().
		Text()
}

func (d MobileBGDecoder) GetOfferImage(doc *goquery.Document) string {
	link, _ := doc.
		Find("tbody tr td.algcent.valgmid a.photoLink img.noborder").
		First().
		Attr("src")

	return link
}

func (d MobileBGDecoder) GetOfferID(doc *goquery.Document) string {
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

func (d MobileBGDecoder) GetOfferLink(doc *goquery.Document) string {
	link, _ := doc.
		Find("tbody tr td.algcent.valgmid a.photoLink").
		First().
		Attr("href")

	return link
}

func (d MobileBGDecoder) GetCarsFromPageResults(pageResults string) map[string]models.CarDTO {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(pageResults))
	var result map[string]models.CarDTO
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
			result[carId] = models.CarDTO{
				ID:          d.GetOfferID(carDoc),
				Title:       d.GetOfferTitle(carDoc),
				Image:       d.GetOfferImage(carDoc),
				Description: d.GetOfferDescription(carDoc),
				Price:       d.GetOfferPrice(carDoc),
				TopOffer:    d.IsTopOffer(carDoc),
			}
		})

	return result
}