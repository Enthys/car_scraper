package mobile_bg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGettingSearchResults(t *testing.T) {
	pageContent := getSearchResults(PageSearchOptions{
		VehicleType: "1",
	})
	t.Log(pageContent)
	assert.Contains(t, pageContent, "<td class=\"valgtop algright\" style=\"width:106px;height:40px\">")
}

func TestParsingSearchOptionsToURLString(t *testing.T) {
	urlString := ParseSearchOptionsToValues(PageSearchOptions{
		Brand: "3",
		VehicleType: "33",
	})

	assert.Equal(t, "act=3&f5=3&pubtype=33", urlString.Encode())
}
