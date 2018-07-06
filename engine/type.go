package engine

import (
	"github.com/PuerkitoBio/goquery"
)

//Request 解析任务
type Request struct {
	URL       string
	ParseFunc func(doc *goquery.Document) (req []Request, err error)
}
