package mobile_bg

import (
	"github.com/PuerkitoBio/goquery"
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
