package mobile_bg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGettingSearchResults(t *testing.T) {
	pageContent, slink := getSearchResults(PageSearchOptions{
		VehicleType: "1",
	})

	assert.Contains(t, pageContent, "<td class=\"valgtop algright\" style=\"width:106px;height:40px\">")
	assert.True(t, len(slink) > 0)
}

func TestParsingSearchOptionsToURLString(t *testing.T) {
	urlString := ParseSearchOptionsToValues(PageSearchOptions{
		Brand: "3",
		VehicleType: "33",
	})

	assert.Equal(t, "act=3&f5=3&pubtype=33", urlString.Encode())
}

func TestRetrievingResultsBySling(t *testing.T) {
	_, slink := getSearchResults(PageSearchOptions{
		VehicleType: "1",
	})

	pageResults := GetSearchBySlink(slink, 2)
	assert.True(t, len(pageResults) > 400)
	assert.Contains(t, pageResults, "<span class=\"pageNumbersSelect\">2</span>")
}