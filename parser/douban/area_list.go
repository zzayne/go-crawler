package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/zzayne/zayne-crawler/engine"
)

//AreaListParser 活动列表解析器
func AreaListParser(doc *goquery.Document) (req []engine.Request, err error) {
	var reqList []engine.Request

	doc.Find(".location a").Each(func(i int, s *goquery.Selection) {
		if url, ok := s.Attr("href"); ok == true {
			reqList = append(reqList, engine.Request{
				URL:       url,
				ParseFunc: CityListParser,
			})
		}
	})

	return reqList, err
}
