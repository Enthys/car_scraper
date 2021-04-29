package mobile_bg

import (
	"car_scraper/models"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func MobileBGGetOfferTitle(doc *goquery.Document) string {
	return doc.
		Find("tbody tr td.valgtop a.mmm").
		First().
		Text()
}

func MobileBGIsTopOffer(doc *goquery.Document) bool {
	return doc.Find("tbody tr td.algright img.noborder").Length() == 3
}

func MobileBGGetOfferDescription(doc *goquery.Document) string {
	return doc.
		Find("tbody tr:nth-child(3) td").
		First().
		Text()
}

func MobileBGGetOfferPrice(doc *goquery.Document) string {
	return doc.
		Find("tbody tr td.algright.valgtop span.price").
		First().
		Text()
}

func MobileBGGetOfferImage(doc *goquery.Document) string {
	link, _ := doc.
		Find("tbody tr td.algcent.valgmid a.photoLink img.noborder").
		First().
		Attr("src")

	return link
}

func MobileBGGetOfferLink(doc *goquery.Document) string {
	link, _ := doc.
		Find("tbody tr td.algcent.valgmid a.photoLink").
		First().
		Attr("href")

	return link
}

func MobileBGGetCarsFromPageResults(pageResults string) []models.CarDTO {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(pageResults))
	var result []models.CarDTO
	doc.Find("form[name=\"search\"]").
		First().
		Find("table").
		Each(func(i int, carTable *goquery.Selection) {
			carDoc := goquery.NewDocumentFromNode(carTable.Nodes[0])
			title := MobileBGGetOfferTitle(carDoc)

			if len(title) == 0 {
				return
			}

			result = append(result, models.CarDTO{
				Title:       MobileBGGetOfferTitle(carDoc),
				Image:       MobileBGGetOfferImage(carDoc),
				Description: MobileBGGetOfferDescription(carDoc),
				Price:       MobileBGGetOfferPrice(carDoc),
				TopOffer:    MobileBGIsTopOffer(carDoc),
			})
		})

	return result
}
