package fetcher

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var rateLimiter = time.Tick(2000 * time.Millisecond)

// Fetch ...
func Fetch(url string) (doc *goquery.Document, err error) {

	<-rateLimiter

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc, err
}
