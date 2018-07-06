package parser

import (
	"errors"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/zzayne/zayne-crawler/engine"
	"github.com/zzayne/zayne-crawler/model"
)

const pageRe = `https://sz.lianjia.com/zufang/pg([0-9]+)`

//RentParser ...
func RentParser(doc *goquery.Document) (req []engine.Request, err error) {

	doc.Find(".house-lst li .info-panel").Each(func(i int, s *goquery.Selection) {
		var house model.House
		house.Img = s.Find("img").AttrOr("src", "")
		house.Name = s.Find("h2 a").AttrOr("title", "unkonwn")
		house.URL = s.Find("h2 a").AttrOr("href", "unkonwn")
		house.Datetime = s.Find("price-pre").Text()

		house.Region = s.Find(".col-1 .where region").Text()
		house.Zone = s.Find(".col-1 .where zone").Text()
		house.Area = s.Find(".col-1 .where meters").Text()
		house.Intro = s.Find(".other .con a").Text()

		if price, err := strconv.Atoi(s.Find(".price .num").Text()); err == nil {
			house.Price = price
		}

		if view, err := strconv.Atoi(s.Find(".col-2 .num").Text()); err == nil {
			house.View = view
		}
	})
	if doc.Length() > 0 {
		// regx := regexp.MustCompile(pageRe)
		// matches := regx.FindAllSubmatch([]byte(), -1)

	}

	return nil, errors.New("end")
}
