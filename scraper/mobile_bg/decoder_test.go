package mobile_bg

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTopOfferIsFoundOnPage(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(topOfferPage))
	assert.True(t, MobileBGIsTopOffer(doc))
}

func TestTopOfferIsNotFoundOnPage(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(normalOfferPage))
	assert.False(t, MobileBGIsTopOffer(doc))
}

func TestRetrievingOfferTitleFromTopOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(topOfferPage))
	assert.Equal(t, "BMW X6 Facelift", MobileBGGetOfferTitle(doc))
}

func TestRetrievingOfferTitleFromNormalOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(normalOfferPage))
	assert.Equal(t, "BMW 114", MobileBGGetOfferTitle(doc))
}

func TestRetrievingDescriptionFromTopOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(topOfferPage))
	assert.Contains(t, MobileBGGetOfferDescription(doc), "Колата е внос от Германия")
}

func TestRetrievingDescriptionFromNormalOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(normalOfferPage))
	assert.Contains(t, MobileBGGetOfferDescription(doc), "инен, Добре запазен автомобил с доказуем пробег.")
}

func TestRetrievingPriceFromTopOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(topOfferPage))
	assert.Contains(t, MobileBGGetOfferPrice(doc), "29 999 лв.")
}

func TestRetrievingPriceFromNormalOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(normalOfferPage))
	assert.Contains(t, MobileBGGetOfferPrice(doc), "15 000 лв.")
}

func TestRetrievingImageLinkFromTopOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(topOfferPage))
	assert.Equal(t, "//mobistatic1.focus.bg/mobile/photosmob/053/2/21611567699623053_nO.jpg", MobileBGGetOfferImage(doc))
}

func TestRetrievingImageLinkFromNormalOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(normalOfferPage))
	assert.Equal(t, "//mobistatic2.focus.bg/mobile/photosmob/455/1/11619071592250455_V9.jpg", MobileBGGetOfferImage(doc))
}

func TestRetrievingLinkFromTopOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(topOfferPage))
	assert.Equal(t, "//www.mobile.bg/pcgi/mobile.cgi?act=4&adv=21611567699623053&slink=jumf9y", MobileBGGetOfferLink(doc))
}

func TestRetrievingLinkFromNormalOffer(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(normalOfferPage))
	assert.Equal(t, "//www.mobile.bg/pcgi/mobile.cgi?act=4&adv=11619071592250455&slink=jumf9y", MobileBGGetOfferLink(doc))
}

const topOfferPage = "<table class=\"tablereset\" style=\"width:660px; margin-bottom:0px; border-top:#008FC6 1px solid; background:url(//www.mobile.bg/images/picturess/top_bg.gif); background-position:bottom; background-repeat:repeat-x;\"><tbody><tr><td style=\"width:220px;height:10px;\"></td><td style=\"width:162px;height:10px;\"></td><td style=\"width:135px;height:10px;\"></td><td style=\"width:37px;height:10px;\"></td><td style=\"width:106px;height:10px;\"></td></tr><tr><td rowspan=\"2\" style=\"width:220px;height:150px\"><table class=\"tablereset\" style=\"width:212px\"><tbody><tr><td class=\"algcent valgmid\"><a href=\"//www.mobile.bg/pcgi/mobile.cgi?act=4&amp;adv=21611567699623053&amp;slink=jumf9y\" class=\"photoLink\"><img src=\"//mobistatic1.focus.bg/mobile/photosmob/053/2/21611567699623053_nO.jpg\" style=\"object-fit: cover;\" class=\"noborder\" alt=\"Обява за продажба на BMW X6 Facelift ~29 999 лв.\" data-geo=\"\" width=\"200\" height=\"150\"></a></td></tr></tbody></table></td><td class=\"valgtop\" style=\"width:162px;height:40px;padding-left:4px\"><a href=\"//www.mobile.bg/pcgi/mobile.cgi?act=4&amp;adv=21611567699623053&amp;slink=jumf9y\" class=\"mmm\">BMW X6 Facelift</a><br><img src=\"//www.mobile.bg/images/picturess/no.gif\" class=\"noborder\" alt=\"\" width=\"1\" height=\"15\"><span style=\"font-size:10px; color:FF0000;\">/нова обява/</span></td><td class=\"algright valgtop\" style=\"width:135px;height:40px;padding-left:4px\"><img src=\"//www.mobile.bg/images/picturess/price-down.png\" style=\"margin-right:3px;\"><span class=\"price\">29 999 лв.</span>&nbsp;</td><td class=\"valgtop\" style=\"width:37px;height:40px\"><a href=\"javascript:;\" id=\"star_21611567699623053\" onclick=\"javascript:openLogPopup(1); return false;\" title=\"Добави обявата в бележника. Изисква регистрация.\" class=\"favListItem\"></a></td><td class=\"valgtop algright\" style=\"width:106px;height:40px\"><a href=\"//cardelux.mobile.bg\" class=\"logoLink\"><img src=\"//cdn2.focus.bg/mobile/images/houseslogos/h14224416982575394.pic?16192639499\" class=\"logoHouse\" alt=\"Лого\"></a></td></tr><tr><td colspan=\"3\" style=\"width:334px;height:50px;padding-left:4px\">дата на произв. - декември 2010 г., пробег - 213000 км, цвят - Черен, Промоционална цена. Колата е внос от Германия, с п...<br>Особености - 4(5) Врати, 4x4, Auto Start Stop function, Bl...<br>Регион: Пловдив, гр. Пловдив</td><td style=\"width:106px\" class=\"algright\"><img src=\"//www.mobile.bg/images/picturess/icons/top.svg\" class=\"noborder\" alt=\"top\" width=\"42\">&nbsp;</td></tr><tr><td colspan=\"5\" style=\"height:5px;\"></td></tr><tr><td colspan=\"2\" style=\"padding-left:4px\"><a href=\"//www.mobile.bg/pcgi/mobile.cgi?act=4&amp;adv=21611567699623053&amp;slink=jumf9y\" class=\"mmm1\">Повече детайли и 15 снимки</a> | <a href=\"javascript:;\" id=\"notepad_21611567699623053\" onclick=\"javascript:openLogPopup(1); return false;\" title=\"Добави обявата в бележника. Изисква регистрация.\" class=\"mmm1\">Добави в бележника</a></td><td colspan=\"3\" class=\"algright\"><a href=\"javascript:;\" id=\"mark_p21611567699623053\" onclick=\"javascript:mark('mark_p21611567699623053',p21611567699623053); updatecomprint('p21611567699623053','//mobistatic1.focus.bg/mobile/photosmob/053/2/med/21611567699623053_nO.jpg');\" class=\"mmm1\">Маркирай обявата</a><img name=\"p21611567699623053\" src=\"//www.mobile.bg/images/picturess/print_n.gif\" class=\"noborder\" alt=\"МАРКИРАЙ ОБЯВАТА\" onclick=\"javascript:mark('mark_p21611567699623053',p21611567699623053); updatecomprint('p21611567699623053','//mobistatic1.focus.bg/mobile/photosmob/053/2/med/21611567699623053_nO.jpg');\" width=\"15\" height=\"15\"><img src=\"//www.mobile.bg/images/picturess/no.gif\" class=\"noborder\" alt=\"\" width=\"4\" height=\"1\"></td></tr><tr><td colspan=\"5\" style=\"height:10px;\"></td></tr></tbody></table>"

const normalOfferPage = "<table class=\"tablereset\" style=\"width:660px; margin-bottom:0px; border-top:#008FC6 1px solid;\"><tbody><tr><td style=\"width:220px;height:10px;\"></td><td style=\"width:162px;height:10px;\"></td><td style=\"width:135px;height:10px;\"></td><td style=\"width:37px;height:10px;\"></td><td style=\"width:106px;height:10px;\"></td></tr><tr><td rowspan=\"2\" style=\"width:220px;height:150px\"><table class=\"tablereset\" style=\"width:212px\"><tbody><tr><td class=\"algcent valgmid\"><a href=\"//www.mobile.bg/pcgi/mobile.cgi?act=4&amp;adv=11619071592250455&amp;slink=jumf9y\" class=\"photoLink\"><img src=\"//mobistatic2.focus.bg/mobile/photosmob/455/1/11619071592250455_V9.jpg\" style=\"object-fit: cover;\" class=\"noborder\" alt=\"Обява за продажба на BMW 114 ~15 000 лв.\" data-geo=\"\" width=\"200\" height=\"150\"></a></td></tr></tbody></table></td><td class=\"valgtop\" style=\"width:162px;height:40px;padding-left:4px\"><a href=\"//www.mobile.bg/pcgi/mobile.cgi?act=4&amp;adv=11619071592250455&amp;slink=jumf9y\" class=\"mmm\">BMW 114</a></td><td class=\"algright valgtop\" style=\"width:135px;height:40px;padding-left:4px\"><span class=\"price\">15 000 лв.</span>&nbsp;</td><td class=\"valgtop\" style=\"width:37px;height:40px\"><a href=\"javascript:;\" id=\"star_11619071592250455\" onclick=\"javascript:openLogPopup(1); return false;\" title=\"Добави обявата в бележника. Изисква регистрация.\" class=\"favListItem\"></a></td><td class=\"valgtop algright\" style=\"width:106px;height:40px\">&nbsp;</td></tr><tr><td colspan=\"4\" style=\"width:440px;height:50px;padding-left:4px\">дата на произв. - април 2014 г., пробег - 93000 км, цвят - Винен, Добре запазен автомобил с доказуем пробег. Без заб...<br>Особености - 4(5) Врати, Bluetooth \\ handsfree система, DV...<br>Регион: София, гр. София</td></tr><tr><td colspan=\"5\" style=\"height:5px;\"></td></tr><tr><td colspan=\"2\" style=\"padding-left:4px\"><a href=\"//www.mobile.bg/pcgi/mobile.cgi?act=4&amp;adv=11619071592250455&amp;slink=jumf9y\" class=\"mmm1\">Повече детайли и 13 снимки</a> | <a href=\"javascript:;\" id=\"notepad_11619071592250455\" onclick=\"javascript:openLogPopup(1); return false;\" title=\"Добави обявата в бележника. Изисква регистрация.\" class=\"mmm1\">Добави в бележника</a></td><td colspan=\"3\" class=\"algright\"><a href=\"javascript:;\" id=\"mark_p11619071592250455\" onclick=\"javascript:mark('mark_p11619071592250455',p11619071592250455); updatecomprint('p11619071592250455','//mobistatic2.focus.bg/mobile/photosmob/455/1/med/11619071592250455_V9.jpg');\" class=\"mmm1\">Маркирай обявата</a><img name=\"p11619071592250455\" src=\"//www.mobile.bg/images/picturess/print_n.gif\" class=\"noborder\" alt=\"МАРКИРАЙ ОБЯВАТА\" onclick=\"javascript:mark('mark_p11619071592250455',p11619071592250455); updatecomprint('p11619071592250455','//mobistatic2.focus.bg/mobile/photosmob/455/1/med/11619071592250455_V9.jpg');\" width=\"15\" height=\"15\"><img src=\"//www.mobile.bg/images/picturess/no.gif\" class=\"noborder\" alt=\"\" width=\"4\" height=\"1\"></td></tr><tr><td colspan=\"5\" style=\"height:10px;\"></td></tr></tbody></table>"
