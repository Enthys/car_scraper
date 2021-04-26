package mobile_bg

import (
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"log"
	"testing"
)

func TestGettingSlink(t *testing.T) {
	resp := getSlink(t)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	out, _ := charmap.Windows1251.NewDecoder().String(string(body))

	t.Logf("HEREERERERRER %s", out)

	//for name, values := range resp.Header {
	//	for _, value := range values {
	//		t.Logf("%v = %v", name, value)
	//	}
	//}
}