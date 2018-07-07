package engine

import (
	"github.com/PuerkitoBio/goquery"
)

//Request 解析任务
type Request struct {
	URL       string
	ParseFunc func(doc *goquery.Document) (RequestResult, error)
}

//RequestResult 解析结果
type RequestResult struct {
	Requests []Request
	Items    []interface{}
}
