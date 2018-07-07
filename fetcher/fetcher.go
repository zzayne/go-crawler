package fetcher

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var rateLimiter = time.Tick(200 * time.Millisecond)

// Fetch 获取对应url的内容，输出文档供解析器解析
func Fetch(url string) (doc *goquery.Document, err error) {
	//设定请求间隔，简单的应对网站反爬虫措施
	<-rateLimiter

	request, err := http.NewRequest(http.MethodGet, url, nil)
	//设置请求header头，简单的应对网站反爬虫措施
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
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc, err
}
