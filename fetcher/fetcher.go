package fetcher

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var rateLimiter = time.Tick(1 * time.Microsecond)

// Fetch ...
func Fetch(url string) (doc *goquery.Document, err error) {

	<-rateLimiter

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//fmt.Println("Redirect:", req)
			return nil
		},
	}

	res, err := client.Do(request)
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
