package parser

import (
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/zzayne/zayne-crawler/engine"
	"github.com/zzayne/zayne-crawler/model"
)

//ActivityParser 活动解析
func ActivityParser(doc *goquery.Document) (req []engine.Request, err error) {

	doc.Find(".picked-events li info").Each(func(i int, s *goquery.Selection) {
		var activity model.Activity
		activity.Name = s.Find("title a").AttrOr("title", "unkonwn")
		activity.URL = s.Find("title a").AttrOr("href", "unkonwn")
		activity.Datetime = s.Find("month").Text() + s.Find("day").Text() + s.Find("time").Text()
		activity.Address = s.Find("address").AttrOr("title", "unkonwn")
		fmt.Println(activity)
	})

	return nil, errors.New("end")
}
