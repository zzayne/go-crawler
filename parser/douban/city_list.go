package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/zzayne/zayne-crawler/engine"
)

//CityListParser 城市列表解析
func CityListParser(doc *goquery.Document) (req []engine.Request, err error) {
	var reqList []engine.Request
	doc.Find(".location a").Each(func(i int, s *goquery.Selection) {
		if url, ok := s.Attr("href"); ok == true {
			reqList = append(reqList, engine.Request{
				URL:       url,
				ParseFunc: ActivityParser,
			})
		}
	})
	return reqList, err
}
