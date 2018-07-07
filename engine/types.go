package engine

import (
	"github.com/PuerkitoBio/goquery"
)

//Request 解析任务
type Request struct {
	URL       string
	ParseFunc func(doc *goquery.Document) (ParseResult, error)
}

//RequestResult 解析结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//Item 数据实体
type Item struct {
	ID      string
	URL     string
	Payload interface{}
}
