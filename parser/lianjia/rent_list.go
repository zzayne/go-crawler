package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/zzayne/go-crawler/engine"
)

//RentListParser ...
func RentListParser(doc *goquery.Document) (engine.ParseResult, error) {
	var result engine.ParseResult
	var URL string

	doc.Find(".house-lst li").Each(func(i int, s *goquery.Selection) {
		URL = s.Find("h2 a").AttrOr("href", "")
		if URL != "" {
			result.Requests = append(result.Requests, engine.Request{
				URL:       URL,
				ParseFunc: RentParser,
			})
		}

	})

	return result, nil
}
