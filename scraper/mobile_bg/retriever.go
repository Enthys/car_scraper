package mobile_bg

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func getSlink(t *testing.T) *http.Response {
	myURL := "https://mobile.bg/pcgi/mobile.cgi"
	nextURL := myURL
	var i int
	var resp *http.Response
	for i < 100 {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			} }

		resp, err := client.Post(
			nextURL,
			"application/x-www-form-urlencoded",
			strings.NewReader(url.Values{"act": {"3"}, "rub": {"1"}, "pubtype": {"1"}}.Encode()),
		)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 200 {
			t.Logf("200 OK: %v", resp)
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

	return resp
}
